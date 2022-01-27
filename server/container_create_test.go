package server_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	types "k8s.io/cri-api/pkg/apis/runtime/v1"
)

// The actual test suite
var _ = t.Describe("ContainerCreate", func() {
	// Prepare the sut
	BeforeEach(func() {
		beforeEach()
		setupSUT()
	})

	AfterEach(afterEach)

	newContainerConfig := func() *types.ContainerConfig {
		return &types.ContainerConfig{
			Metadata: &types.ContainerMetadata{},
			Image:    &types.ImageSpec{},
			Linux: &types.LinuxContainerConfig{
				Resources: &types.LinuxContainerResources{},
				SecurityContext: &types.LinuxContainerSecurityContext{
					Capabilities:     &types.Capability{},
					NamespaceOptions: &types.NamespaceOption{},
					SelinuxOptions:   &types.SELinuxOption{},
					RunAsUser:        &types.Int64Value{},
					RunAsGroup:       &types.Int64Value{},
				},
			},
		}
	}

	newPodSandboxConfig := func() *types.PodSandboxConfig {
		return &types.PodSandboxConfig{
			Metadata:     &types.PodSandboxMetadata{},
			DnsConfig:    &types.DNSConfig{},
			PortMappings: []*types.PortMapping{},
			Linux: &types.LinuxPodSandboxConfig{
				SecurityContext: &types.LinuxSandboxSecurityContext{
					NamespaceOptions: &types.NamespaceOption{},
					SelinuxOptions:   &types.SELinuxOption{},
					RunAsUser:        &types.Int64Value{},
					RunAsGroup:       &types.Int64Value{},
				},
			},
		}
	}

	t.Describe("ContainerCreate", func() {
		It("should fail when container config image is nil", func() {
			// Given
			addContainerAndSandbox()

			// When
			response, err := sut.CreateContainer(context.Background(),
				&types.CreateContainerRequest{
					PodSandboxId: testSandbox.ID(),
					Config: &types.ContainerConfig{
						Metadata: &types.ContainerMetadata{
							Name: "name",
						},
					},
				})

			// Then
			Expect(err).NotTo(BeNil())
			Expect(response).To(BeNil())
		})

		It("should fail when container config metadata name is empty", func() {
			// Given
			addContainerAndSandbox()

			// When
			response, err := sut.CreateContainer(context.Background(),
				&types.CreateContainerRequest{
					PodSandboxId: testSandbox.ID(),
					Config: &types.ContainerConfig{
						Metadata: &types.ContainerMetadata{},
					},
				})

			// Then
			Expect(err).NotTo(BeNil())
			Expect(response).To(BeNil())
		})

		It("should fail when container config metadata is nil", func() {
			// Given
			addContainerAndSandbox()

			// When
			response, err := sut.CreateContainer(context.Background(),
				&types.CreateContainerRequest{
					PodSandboxId: testSandbox.ID(),
					Config:       &types.ContainerConfig{},
				})

			// Then
			Expect(err).NotTo(BeNil())
			Expect(response).To(BeNil())
		})

		It("should fail when container config is nil", func() {
			// Given
			addContainerAndSandbox()

			// When
			response, err := sut.CreateContainer(context.Background(),
				&types.CreateContainerRequest{
					PodSandboxId:  testSandbox.ID(),
					SandboxConfig: newPodSandboxConfig(),
				})

			// Then
			Expect(err).NotTo(BeNil())
			Expect(response).To(BeNil())
		})

		It("should fail when container is stopped", func() {
			// Given
			addContainerAndSandbox()
			testSandbox.SetStopped(false)

			// When
			response, err := sut.CreateContainer(context.Background(),
				&types.CreateContainerRequest{
					PodSandboxId:  testSandbox.ID(),
					Config:        newContainerConfig(),
					SandboxConfig: newPodSandboxConfig(),
				})

			// Then
			Expect(err).NotTo(BeNil())
			Expect(response).To(BeNil())
		})

		It("should fail when sandbox not found", func() {
			// Given
			Expect(sut.PodIDIndex().Add(testSandbox.ID())).To(BeNil())

			// When
			response, err := sut.CreateContainer(context.Background(),
				&types.CreateContainerRequest{
					PodSandboxId:  testSandbox.ID(),
					Config:        newContainerConfig(),
					SandboxConfig: newPodSandboxConfig(),
				})

			// Then
			Expect(err).NotTo(BeNil())
			Expect(response).To(BeNil())
		})

		It("should fail on invalid pod sandbox ID", func() {
			// Given
			// When
			response, err := sut.CreateContainer(context.Background(),
				&types.CreateContainerRequest{
					PodSandboxId:  testSandbox.ID(),
					Config:        newContainerConfig(),
					SandboxConfig: newPodSandboxConfig(),
				})

			// Then
			Expect(err).NotTo(BeNil())
			Expect(response).To(BeNil())
		})

		It("should fail on empty pod sandbox ID", func() {
			// Given
			// When
			response, err := sut.CreateContainer(context.Background(),
				&types.CreateContainerRequest{
					Config:        newContainerConfig(),
					SandboxConfig: newPodSandboxConfig(),
				})

			// Then
			Expect(err).NotTo(BeNil())
			Expect(response).To(BeNil())
		})
	})
})
