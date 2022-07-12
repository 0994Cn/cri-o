# CRI-O Dependency Report

_Generated on Tue, 12 Jul 2022 22:26:19 UTC for commit [07433d7][0]._

[0]: https://github.com/cri-o/cri-o/commit/07433d7cb0a149c0cd5bc64d0080e26d20ba6591

## Outdated Dependencies

|                                   MODULE                                    |                    VERSION                    |            NEW VERSION             | DIRECT | VALID TIMESTAMPS |
|-----------------------------------------------------------------------------|-----------------------------------------------|------------------------------------|--------|------------------|
| github.com/containers/conmon-rs                                             | v0.0.0-20220609131637-836ba11bf0aa            | v0.0.0-20220712072830-20eb3b35a5cd | true   | true             |
| github.com/emicklei/go-restful                                              | v2.15.0+incompatible                          | v2.16.0+incompatible               | true   | true             |
| github.com/urfave/cli/v2                                                    | v2.10.3                                       | v2.11.0                            | true   | true             |
| go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc | v0.32.0                                       | v0.33.0                            | true   | true             |
| go.opentelemetry.io/otel                                                    | v1.7.0                                        | v1.8.0                             | true   | true             |
| go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc             | v1.7.0                                        | v1.8.0                             | true   | true             |
| go.opentelemetry.io/otel/sdk                                                | v1.7.0                                        | v1.8.0                             | true   | true             |
| golang.org/x/net                                                            | v0.0.0-20220621193019-9d032be2e588            | v0.0.0-20220708220712-1185a9018129 | true   | true             |
| golang.org/x/sys                                                            | v0.0.0-20220615213510-4f61da869c0c            | v0.0.0-20220712014510-0a85c31ab51e | true   | true             |
| google.golang.org/grpc                                                      | v1.47.0                                       | v1.48.0                            | true   | true             |
| k8s.io/api                                                                  | v0.0.0-20220624161656-6219eed24f3c            | v0.0.0-20220712183556-71481bf24746 | true   | true             |
| k8s.io/apimachinery                                                         | v0.0.0-20220624161656-6219eed24f3c            | v0.0.0-20220712183556-71481bf24746 | true   | true             |
| k8s.io/cri-api                                                              | v0.0.0-20220624161656-6219eed24f3c            | v0.0.0-20220712183556-71481bf24746 | true   | true             |
| k8s.io/kubernetes                                                           | v1.25.0-alpha.1.0.20220624161656-6219eed24f3c | v1.25.0-alpha.2                    | true   | true             |
| k8s.io/utils                                                                | v0.0.0-20220210201930-3a6ce19ff2f9            | v0.0.0-20220706174534-f6158b442e7c | true   | true             |
| sigs.k8s.io/bom                                                             | v0.3.0-rc1.0.20220627190259-3850e8ff14fa      | v0.3.0                             | true   | true             |

## All Dependencies

|                                   MODULE                                    |                      VERSION                      |            NEW VERSION             | DIRECT | VALID TIMESTAMPS |
|-----------------------------------------------------------------------------|---------------------------------------------------|------------------------------------|--------|------------------|
| 4d63.com/gochecknoglobals                                                   | v0.1.0                                            |                                    | false  | true             |
| bazil.org/fuse                                                              | v0.0.0-20200407214033-5883e5a4b512                | v0.0.0-20200524192727-fb710f7dfd05 | false  | true             |
| bitbucket.org/bertimus9/systemstat                                          | v0.5.0                                            |                                    | false  | true             |
| bitbucket.org/creachadair/shell                                             | v0.0.6                                            | v0.0.7                             | false  | true             |
| bou.ke/monkey                                                               | v1.0.2                                            |                                    | false  | true             |
| capnproto.org/go/capnp/v3                                                   | v3.0.0-alpha.3                                    |                                    | false  | true             |
| cloud.google.com/go                                                         | v0.100.2                                          | v0.103.0                           | false  | true             |
| cloud.google.com/go/bigquery                                                | v1.8.0                                            | v1.35.0                            | false  | true             |
| cloud.google.com/go/compute                                                 | v1.6.1                                            | v1.7.0                             | false  | true             |
| cloud.google.com/go/containeranalysis                                       | v0.3.0                                            | v0.4.0                             | false  | true             |
| cloud.google.com/go/datastore                                               | v1.1.0                                            | v1.8.0                             | false  | true             |
| cloud.google.com/go/errorreporting                                          | v0.2.0                                            |                                    | false  | true             |
| cloud.google.com/go/firestore                                               | v1.6.1                                            |                                    | false  | true             |
| cloud.google.com/go/grafeas                                                 | v0.1.0                                            | v0.2.0                             | false  | true             |
| cloud.google.com/go/iam                                                     | v0.3.0                                            |                                    | false  | true             |
| cloud.google.com/go/kms                                                     | v1.4.0                                            |                                    | false  | true             |
| cloud.google.com/go/logging                                                 | v1.4.2                                            | v1.5.0                             | false  | true             |
| cloud.google.com/go/monitoring                                              | v1.1.0                                            | v1.5.0                             | false  | true             |
| cloud.google.com/go/pubsub                                                  | v1.17.1                                           | v1.23.1                            | false  | true             |
| cloud.google.com/go/secretmanager                                           | v1.0.0                                            | v1.5.0                             | false  | true             |
| cloud.google.com/go/security                                                | v1.1.1                                            | v1.4.1                             | false  | true             |
| cloud.google.com/go/spanner                                                 | v1.25.0                                           | v1.34.1                            | false  | true             |
| cloud.google.com/go/storage                                                 | v1.22.1                                           | v1.23.0                            | false  | true             |
| cloud.google.com/go/trace                                                   | v1.0.0                                            | v1.2.0                             | false  | true             |
| code.gitea.io/sdk/gitea                                                     | v0.11.3                                           | v0.15.1                            | false  | true             |
| contrib.go.opencensus.io/exporter/aws                                       | v0.0.0-20200617204711-c478e41e60e9                |                                    | false  | true             |
| contrib.go.opencensus.io/exporter/ocagent                                   | v0.7.1-0.20200907061046-05415f1de66d              |                                    | false  | true             |
| contrib.go.opencensus.io/exporter/prometheus                                | v0.4.0                                            | v0.4.1                             | false  | true             |
| contrib.go.opencensus.io/exporter/stackdriver                               | v0.13.10                                          | v0.13.13                           | false  | true             |
| contrib.go.opencensus.io/exporter/zipkin                                    | v0.1.2                                            |                                    | false  | true             |
| contrib.go.opencensus.io/integrations/ocsql                                 | v0.1.7                                            |                                    | false  | true             |
| contrib.go.opencensus.io/resource                                           | v0.1.1                                            | v0.1.2                             | false  | true             |
| cuelang.org/go                                                              | v0.4.3                                            |                                    | false  | true             |
| dmitri.shuralyov.com/gpu/mtl                                                | v0.0.0-20201218220906-28db891af037                |                                    | false  | true             |
| filippo.io/edwards25519                                                     | v1.0.0-rc.1                                       | v1.0.0                             | false  | true             |
| github.com/14rcole/gopopulate                                               | v0.0.0-20180821133914-b175b219e774                |                                    | false  | true             |
| github.com/AdaLogics/go-fuzz-headers                                        | v0.0.0-20211102141018-f7be0cbad29c                | v0.0.0-20220708163326-82d177caec6e | false  | true             |
| github.com/Antonboom/errname                                                | v0.1.5                                            | v0.1.7                             | false  | true             |
| github.com/Antonboom/nilnil                                                 | v0.1.0                                            | v0.1.1                             | false  | true             |
| github.com/Azure/azure-amqp-common-go/v2                                    | v2.1.0                                            | v2.1.1                             | false  | true             |
| github.com/Azure/azure-amqp-common-go/v3                                    | v3.2.2                                            | v3.2.3                             | false  | true             |
| github.com/Azure/azure-pipeline-go                                          | v0.2.3                                            |                                    | false  | true             |
| github.com/Azure/azure-sdk-for-go                                           | v63.3.0+incompatible                              | v66.0.0+incompatible               | false  | true             |
| github.com/Azure/azure-service-bus-go                                       | v0.11.5                                           |                                    | false  | true             |
| github.com/Azure/azure-storage-blob-go                                      | v0.14.0                                           | v0.15.0                            | false  | true             |
| github.com/Azure/go-amqp                                                    | v0.16.4                                           | v0.17.5                            | false  | true             |
| github.com/Azure/go-ansiterm                                                | v0.0.0-20210617225240-d185dfc1b5a1                |                                    | false  | true             |
| github.com/Azure/go-autorest                                                | v14.2.0+incompatible                              |                                    | false  | true             |
| github.com/Azure/go-autorest/autorest                                       | v0.11.27                                          |                                    | false  | true             |
| github.com/Azure/go-autorest/autorest/adal                                  | v0.9.20                                           |                                    | false  | true             |
| github.com/Azure/go-autorest/autorest/azure/auth                            | v0.5.11                                           |                                    | false  | true             |
| github.com/Azure/go-autorest/autorest/azure/cli                             | v0.4.5                                            |                                    | false  | true             |
| github.com/Azure/go-autorest/autorest/date                                  | v0.3.0                                            |                                    | false  | true             |
| github.com/Azure/go-autorest/autorest/mocks                                 | v0.4.2                                            |                                    | false  | true             |
| github.com/Azure/go-autorest/autorest/to                                    | v0.4.0                                            |                                    | false  | true             |
| github.com/Azure/go-autorest/autorest/validation                            | v0.3.1                                            |                                    | false  | true             |
| github.com/Azure/go-autorest/logger                                         | v0.2.1                                            |                                    | false  | true             |
| github.com/Azure/go-autorest/tracing                                        | v0.6.0                                            |                                    | false  | true             |
| github.com/BurntSushi/toml                                                  | v1.1.0                                            |                                    | true   | true             |
| github.com/BurntSushi/xgb                                                   | v0.0.0-20160522181843-27f122750802                | v0.0.0-20210121224620-deaf085860bc | false  | true             |
| github.com/DataDog/datadog-go                                               | v3.2.0+incompatible                               | v4.8.3+incompatible                | false  | true             |
| github.com/Djarvur/go-err113                                                | v0.0.0-20210108212216-aea10b59be24                |                                    | false  | true             |
| github.com/GoogleCloudPlatform/cloudsql-proxy                               | v1.27.0                                           | v1.31.1                            | false  | true             |
| github.com/GoogleCloudPlatform/k8s-cloud-provider                           | v1.18.1-0.20220218231025-f11817397a1b             |                                    | false  | true             |
| github.com/GoogleCloudPlatform/testgrid                                     | v0.0.38                                           | v0.0.144                           | false  | true             |
| github.com/JeffAshton/win_pdh                                               | v0.0.0-20161109143554-76bb4ee9f0ab                |                                    | false  | true             |
| github.com/Knetic/govaluate                                                 | v3.0.1-0.20171022003610-9aa49832a739+incompatible |                                    | false  | true             |
| github.com/MakeNowJust/heredoc                                              | v1.0.0                                            |                                    | false  | true             |
| github.com/Masterminds/goutils                                              | v1.1.1                                            |                                    | false  | true             |
| github.com/Masterminds/semver                                               | v1.5.0                                            |                                    | false  | true             |
| github.com/Masterminds/semver/v3                                            | v3.1.1                                            |                                    | false  | true             |
| github.com/Masterminds/sprig                                                | v2.22.0+incompatible                              |                                    | false  | true             |
| github.com/Masterminds/sprig/v3                                             | v3.2.2                                            |                                    | false  | true             |
| github.com/Microsoft/go-winio                                               | v0.5.2                                            |                                    | true   | true             |
| github.com/Microsoft/hcsshim                                                | v0.9.3                                            |                                    | false  | true             |
| github.com/Microsoft/hcsshim/test                                           | v0.0.0-20210227013316-43a75bb4edd3                | v0.0.0-20220712165458-94f78da96a60 | false  | true             |
| github.com/NYTimes/gziphandler                                              | v1.1.1                                            |                                    | false  | true             |
| github.com/OneOfOne/xxhash                                                  | v1.2.8                                            |                                    | false  | true             |
| github.com/OpenPeeDeeP/depguard                                             | v1.0.1                                            | v1.1.0                             | false  | true             |
| github.com/PaesslerAG/gval                                                  | v1.0.0                                            | v1.2.0                             | false  | true             |
| github.com/PaesslerAG/jsonpath                                              | v0.1.1                                            |                                    | false  | true             |
| github.com/ProtonMail/go-crypto                                             | v0.0.0-20220407094043-a94812496cf5                | v0.0.0-20220711121315-1fde58898e96 | false  | true             |
| github.com/PuerkitoBio/purell                                               | v1.1.1                                            |                                    | false  | true             |
| github.com/PuerkitoBio/urlesc                                               | v0.0.0-20170810143723-de5bf2ad4578                |                                    | false  | true             |
| github.com/ReneKroon/ttlcache/v2                                            | v2.11.0                                           | v2.11.1                            | false  | true             |
| github.com/Shopify/logrus-bugsnag                                           | v0.0.0-20171204204709-577dee27f20d                |                                    | false  | true             |
| github.com/Shopify/sarama                                                   | v1.19.0                                           | v1.34.1                            | false  | true             |
| github.com/Shopify/toxiproxy                                                | v2.1.4+incompatible                               |                                    | false  | true             |
| github.com/StackExchange/wmi                                                | v1.2.1                                            |                                    | false  | true             |
| github.com/ThalesIgnite/crypto11                                            | v1.2.5                                            |                                    | false  | true             |
| github.com/VividCortex/ewma                                                 | v1.2.0                                            |                                    | false  | true             |
| github.com/VividCortex/gohistogram                                          | v1.0.0                                            |                                    | false  | true             |
| github.com/acarl005/stripansi                                               | v0.0.0-20180116102854-5a71ef0e047d                |                                    | false  | true             |
| github.com/acomagu/bufpipe                                                  | v1.0.3                                            |                                    | false  | true             |
| github.com/afex/hystrix-go                                                  | v0.0.0-20180502004556-fa1af6a1f4f5                |                                    | false  | true             |
| github.com/agnivade/levenshtein                                             | v1.0.1                                            | v1.1.1                             | false  | true             |
| github.com/ajstarks/svgo                                                    | v0.0.0-20180226025133-644b8db467af                | v0.0.0-20211024235047-1546f124cd8b | false  | true             |
| github.com/alcortesm/tgz                                                    | v0.0.0-20161220082320-9c5fe88206d7                |                                    | false  | true             |
| github.com/alecthomas/kingpin                                               | v2.2.6+incompatible                               |                                    | false  | true             |
| github.com/alecthomas/template                                              | v0.0.0-20190718012654-fb15b899a751                |                                    | false  | true             |
| github.com/alecthomas/units                                                 | v0.0.0-20190924025748-f65c72e2690d                | v0.0.0-20211218093645-b94a6e3cc137 | false  | true             |
| github.com/alexflint/go-filemutex                                           | v1.1.0                                            | v1.2.0                             | false  | true             |
| github.com/alexkohler/prealloc                                              | v1.0.0                                            |                                    | false  | true             |
| github.com/andreyvit/diff                                                   | v0.0.0-20170406064948-c7f18ee00883                |                                    | false  | true             |
| github.com/andybalholm/brotli                                               | v1.0.3                                            | v1.0.4                             | false  | true             |
| github.com/anmitsu/go-shlex                                                 | v0.0.0-20200514113438-38f4b401e2be                |                                    | false  | true             |
| github.com/antihax/optional                                                 | v1.0.0                                            |                                    | false  | true             |
| github.com/antlr/antlr4/runtime/Go/antlr                                    | v0.0.0-20220209173558-ad29539cd2e9                | v0.0.0-20220626175859-9abda183db8e | false  | true             |
| github.com/aokoli/goutils                                                   | v1.0.1                                            | v1.1.1                             | false  | true             |
| github.com/apache/beam                                                      | v2.32.0+incompatible                              | v2.40.0+incompatible               | false  | true             |
| github.com/apache/thrift                                                    | v0.13.0                                           | v0.16.0                            | false  | true             |
| github.com/apex/log                                                         | v1.1.4                                            | v1.9.0                             | false  | true             |
| github.com/apex/logs                                                        | v0.0.4                                            | v1.1.0                             | false  | true             |
| github.com/aphistic/golf                                                    | v0.0.0-20180712155816-02c07f170c5a                |                                    | false  | true             |
| github.com/aphistic/sweet                                                   | v0.2.0                                            | v0.3.0                             | false  | true             |
| github.com/armon/circbuf                                                    | v0.0.0-20150827004946-bbbad097214e                | v0.0.0-20190214190532-5111143e8da2 | false  | true             |
| github.com/armon/consul-api                                                 | v0.0.0-20180202201655-eb2c6b5be1b6                |                                    | false  | true             |
| github.com/armon/go-metrics                                                 | v0.4.0                                            |                                    | false  | true             |
| github.com/armon/go-radix                                                   | v1.0.0                                            |                                    | false  | true             |
| github.com/armon/go-socks5                                                  | v0.0.0-20160902184237-e75332964ef5                |                                    | false  | true             |
| github.com/aryann/difflib                                                   | v0.0.0-20170710044230-e206f873d14a                | v0.0.0-20210328193216-ff5ff6dc229b | false  | true             |
| github.com/asaskevich/govalidator                                           | v0.0.0-20210307081110-f21760c49a8d                |                                    | false  | true             |
| github.com/ashanbrown/forbidigo                                             | v1.2.0                                            | v1.3.0                             | false  | true             |
| github.com/ashanbrown/makezero                                              | v0.0.0-20210520155254-b6261585ddde                | v1.1.1                             | false  | true             |
| github.com/auth0/go-jwt-middleware                                          | v1.0.1                                            |                                    | false  | true             |
| github.com/aws/aws-lambda-go                                                | v1.13.3                                           | v1.32.1                            | false  | true             |
| github.com/aws/aws-sdk-go                                                   | v1.43.45                                          | v1.44.53                           | false  | true             |
| github.com/aws/aws-sdk-go-v2                                                | v1.16.4                                           | v1.16.7                            | false  | true             |
| github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream                       | v1.0.0                                            | v1.4.3                             | false  | true             |
| github.com/aws/aws-sdk-go-v2/config                                         | v1.14.0                                           | v1.15.14                           | false  | true             |
| github.com/aws/aws-sdk-go-v2/credentials                                    | v1.9.0                                            | v1.12.9                            | false  | true             |
| github.com/aws/aws-sdk-go-v2/feature/ec2/imds                               | v1.11.0                                           | v1.12.8                            | false  | true             |
| github.com/aws/aws-sdk-go-v2/feature/s3/manager                             | v1.7.1                                            | v1.11.20                           | false  | true             |
| github.com/aws/aws-sdk-go-v2/internal/configsources                         | v1.1.5                                            | v1.1.14                            | false  | true             |
| github.com/aws/aws-sdk-go-v2/internal/endpoints/v2                          | v2.3.0                                            | v2.4.8                             | false  | true             |
| github.com/aws/aws-sdk-go-v2/internal/ini                                   | v1.3.6                                            | v1.3.15                            | false  | true             |
| github.com/aws/aws-sdk-go-v2/service/ecr                                    | v1.15.0                                           | v1.17.8                            | false  | true             |
| github.com/aws/aws-sdk-go-v2/service/ecrpublic                              | v1.12.0                                           | v1.13.8                            | false  | true             |
| github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding               | v1.5.0                                            | v1.9.3                             | false  | true             |
| github.com/aws/aws-sdk-go-v2/service/internal/presigned-url                 | v1.8.0                                            | v1.9.8                             | false  | true             |
| github.com/aws/aws-sdk-go-v2/service/internal/s3shared                      | v1.9.0                                            | v1.13.8                            | false  | true             |
| github.com/aws/aws-sdk-go-v2/service/kms                                    | v1.10.0                                           | v1.17.5                            | false  | true             |
| github.com/aws/aws-sdk-go-v2/service/s3                                     | v1.19.0                                           | v1.27.1                            | false  | true             |
| github.com/aws/aws-sdk-go-v2/service/secretsmanager                         | v1.10.0                                           | v1.15.13                           | false  | true             |
| github.com/aws/aws-sdk-go-v2/service/sns                                    | v1.11.0                                           | v1.17.9                            | false  | true             |
| github.com/aws/aws-sdk-go-v2/service/sqs                                    | v1.12.0                                           | v1.19.0                            | false  | true             |
| github.com/aws/aws-sdk-go-v2/service/ssm                                    | v1.15.0                                           | v1.27.4                            | false  | true             |
| github.com/aws/aws-sdk-go-v2/service/sso                                    | v1.10.0                                           | v1.11.12                           | false  | true             |
| github.com/aws/aws-sdk-go-v2/service/sts                                    | v1.15.0                                           | v1.16.9                            | false  | true             |
| github.com/aws/smithy-go                                                    | v1.11.2                                           | v1.12.0                            | false  | true             |
| github.com/awslabs/amazon-ecr-credential-helper/ecr-login                   | v0.0.0-20220228164355-396b2034c795                | v0.0.0-20220706184558-ce46abcd012b | false  | true             |
| github.com/aybabtme/rgbterm                                                 | v0.0.0-20170906152045-cc83f3b3ce59                |                                    | false  | true             |
| github.com/beeker1121/goque                                                 | v1.0.3-0.20191103205551-d618510128af              |                                    | false  | true             |
| github.com/beevik/etree                                                     | v1.1.0                                            |                                    | false  | true             |
| github.com/benbjohnson/clock                                                | v1.1.0                                            | v1.3.0                             | false  | true             |
| github.com/beorn7/perks                                                     | v1.0.1                                            |                                    | false  | true             |
| github.com/bgentry/speakeasy                                                | v0.1.0                                            |                                    | false  | true             |
| github.com/bitly/go-simplejson                                              | v0.5.0                                            |                                    | false  | true             |
| github.com/bits-and-blooms/bitset                                           | v1.2.0                                            | v1.2.2                             | false  | true             |
| github.com/bketelsen/crypt                                                  | v0.0.4                                            | v0.0.5                             | false  | true             |
| github.com/bkielbasa/cyclop                                                 | v1.2.0                                            |                                    | false  | true             |
| github.com/blakesmith/ar                                                    | v0.0.0-20190502131153-809d4375e1fb                |                                    | false  | true             |
| github.com/blang/semver                                                     | v3.5.1+incompatible                               |                                    | true   | true             |
| github.com/blang/semver/v4                                                  | v4.0.0                                            |                                    | false  | true             |
| github.com/blendle/zapdriver                                                | v1.3.1                                            |                                    | false  | true             |
| github.com/blizzy78/varnamelen                                              | v0.3.0                                            | v0.8.0                             | false  | true             |
| github.com/bmizerany/assert                                                 | v0.0.0-20160611221934-b7ed37b82869                |                                    | false  | true             |
| github.com/boltdb/bolt                                                      | v1.3.1                                            |                                    | false  | true             |
| github.com/bombsimon/wsl/v3                                                 | v3.3.0                                            |                                    | false  | true             |
| github.com/bradfitz/gomemcache                                              | v0.0.0-20190913173617-a41fca850d0b                | v0.0.0-20220106215444-fb4bf637b56d | false  | true             |
| github.com/bradleyfalzon/ghinstallation/v2                                  | v2.0.3                                            | v2.1.0                             | false  | true             |
| github.com/breml/bidichk                                                    | v0.1.1                                            | v0.2.3                             | false  | true             |
| github.com/bshuster-repo/logrus-logstash-hook                               | v0.4.1                                            | v1.0.2                             | false  | true             |
| github.com/buger/goterm                                                     | v1.0.4                                            |                                    | false  | true             |
| github.com/buger/jsonparser                                                 | v1.1.1                                            |                                    | false  | true             |
| github.com/bugsnag/bugsnag-go                                               | v0.0.0-20141110184014-b1d153021fcd                | v2.1.2+incompatible                | false  | true             |
| github.com/bugsnag/osext                                                    | v0.0.0-20130617224835-0dd3f918b21b                |                                    | false  | true             |
| github.com/bugsnag/panicwrap                                                | v0.0.0-20151223152923-e2c28503fcd0                | v1.3.4                             | false  | true             |
| github.com/butuzov/ireturn                                                  | v0.1.1                                            |                                    | false  | true             |
| github.com/bytecodealliance/wasmtime-go                                     | v0.33.1                                           | v0.38.1                            | false  | true             |
| github.com/caarlos0/ctrlc                                                   | v1.0.0                                            | v1.1.0                             | false  | true             |
| github.com/campoy/unique                                                    | v0.0.0-20180121183637-88950e537e7e                |                                    | false  | true             |
| github.com/carolynvs/magex                                                  | v0.8.1                                            |                                    | false  | true             |
| github.com/casbin/casbin/v2                                                 | v2.1.2                                            | v2.50.2                            | false  | true             |
| github.com/cavaliercoder/badio                                              | v0.0.0-20160213150051-ce5280129e9e                |                                    | false  | true             |
| github.com/cavaliercoder/go-cpio                                            | v0.0.0-20180626203310-925f9528c45e                | v1.0.1                             | false  | true             |
| github.com/cavaliercoder/go-rpm                                             | v0.0.0-20200122174316-8cb9fd9c31a8                | v1.2.0                             | false  | true             |
| github.com/cenkalti/backoff                                                 | v2.2.1+incompatible                               |                                    | false  | true             |
| github.com/cenkalti/backoff/v3                                              | v3.2.2                                            |                                    | false  | true             |
| github.com/cenkalti/backoff/v4                                              | v4.1.3                                            |                                    | false  | true             |
| github.com/census-instrumentation/opencensus-proto                          | v0.3.0                                            |                                    | false  | true             |
| github.com/certifi/gocertifi                                                | v0.0.0-20200922220541-2c3bb06c6054                | v0.0.0-20210507211836-431795d63e8d | false  | true             |
| github.com/cespare/xxhash                                                   | v1.1.0                                            |                                    | false  | true             |
| github.com/cespare/xxhash/v2                                                | v2.1.2                                            |                                    | false  | true             |
| github.com/chai2010/gettext-go                                              | v1.0.2                                            |                                    | false  | true             |
| github.com/charithe/durationcheck                                           | v0.0.9                                            |                                    | false  | true             |
| github.com/chavacava/garif                                                  | v0.0.0-20210405164556-e8a0a408d6af                | v0.0.0-20220630083739-93517212f375 | false  | true             |
| github.com/checkpoint-restore/checkpointctl                                 | v0.0.0-20220321135231-33f4a66335f0                |                                    | false  | true             |
| github.com/checkpoint-restore/go-criu/v4                                    | v4.1.0                                            |                                    | false  | true             |
| github.com/checkpoint-restore/go-criu/v5                                    | v5.3.0                                            |                                    | false  | true             |
| github.com/cheggaaa/pb/v3                                                   | v3.0.8                                            | v3.1.0                             | false  | true             |
| github.com/chrismellard/docker-credential-acr-env                           | v0.0.0-20220119192733-fe33c00cee21                | v0.0.0-20220327082430-c57b701bfc08 | false  | true             |
| github.com/chzyer/logex                                                     | v1.1.10                                           | v1.2.1                             | false  | true             |
| github.com/chzyer/readline                                                  | v0.0.0-20180603132655-2972be24d48e                | v1.5.0                             | false  | true             |
| github.com/chzyer/test                                                      | v0.0.0-20180213035817-a1ea475d72b1                | v1.0.0                             | false  | true             |
| github.com/cilium/ebpf                                                      | v0.7.0                                            | v0.9.0                             | false  | true             |
| github.com/circonus-labs/circonus-gometrics                                 | v2.3.1+incompatible                               |                                    | false  | true             |
| github.com/circonus-labs/circonusllhist                                     | v0.1.3                                            | v0.3.0                             | false  | true             |
| github.com/clbanning/x2j                                                    | v0.0.0-20191024224557-825249438eec                |                                    | false  | true             |
| github.com/client9/misspell                                                 | v0.3.4                                            |                                    | false  | true             |
| github.com/clusterhq/flocker-go                                             | v0.0.0-20160920122132-2b8b7259d313                |                                    | false  | true             |
| github.com/cncf/udpa/go                                                     | v0.0.0-20210930031921-04548b0d99d4                | v0.0.0-20220112060539-c52dc94e7fbe | false  | true             |
| github.com/cncf/xds/go                                                      | v0.0.0-20211130200136-a8f946100490                | v0.0.0-20220520190051-1e77728a1eaa | false  | true             |
| github.com/cockroachdb/apd                                                  | v1.1.0                                            |                                    | false  | true             |
| github.com/cockroachdb/apd/v2                                               | v2.0.1                                            | v2.0.2                             | false  | true             |
| github.com/cockroachdb/datadriven                                           | v0.0.0-20200714090401-bf6692d28da5                | v1.0.1                             | false  | true             |
| github.com/cockroachdb/errors                                               | v1.2.4                                            | v1.9.0                             | false  | true             |
| github.com/cockroachdb/logtags                                              | v0.0.0-20190617123548-eb05cc24525f                | v0.0.0-20211118104740-dabe8e521a4f | false  | true             |
| github.com/codahale/hdrhistogram                                            | v0.0.0-20161010025455-3a0bb77429bd                | v1.1.2                             | false  | true             |
| github.com/codahale/rfc6979                                                 | v0.0.0-20141003034818-6a90f24967eb                |                                    | false  | true             |
| github.com/common-nighthawk/go-figure                                       | v0.0.0-20210622060536-734e95fb86be                |                                    | false  | true             |
| github.com/container-orchestrated-devices/container-device-interface        | v0.4.0                                            |                                    | true   | true             |
| github.com/container-storage-interface/spec                                 | v1.6.0                                            |                                    | false  | true             |
| github.com/containerd/aufs                                                  | v1.0.0                                            |                                    | false  | true             |
| github.com/containerd/btrfs                                                 | v1.0.0                                            |                                    | false  | true             |
| github.com/containerd/cgroups                                               | v1.0.4                                            |                                    | true   | true             |
| github.com/containerd/console                                               | v1.0.3                                            |                                    | false  | true             |
| github.com/containerd/containerd                                            | v1.6.6                                            |                                    | true   | true             |
| github.com/containerd/continuity                                            | v0.2.2                                            | v0.3.0                             | false  | true             |
| github.com/containerd/cri-containerd                                        | v1.19.0                                           |                                    | true   | true             |
| github.com/containerd/fifo                                                  | v1.0.0                                            |                                    | true   | true             |
| github.com/containerd/go-cni                                                | v1.1.6                                            |                                    | false  | true             |
| github.com/containerd/go-runc                                               | v1.0.0                                            |                                    | false  | true             |
| github.com/containerd/imgcrypt                                              | v1.1.4                                            | v1.1.6                             | false  | true             |
| github.com/containerd/nri                                                   | v0.1.0                                            |                                    | false  | true             |
| github.com/containerd/stargz-snapshotter/estargz                            | v0.11.4                                           | v0.12.0                            | false  | true             |
| github.com/containerd/ttrpc                                                 | v1.1.0                                            |                                    | true   | true             |
| github.com/containerd/typeurl                                               | v1.0.2                                            |                                    | true   | true             |
| github.com/containerd/zfs                                                   | v1.0.0                                            |                                    | false  | true             |
| github.com/containernetworking/cni                                          | v1.1.1                                            |                                    | true   | true             |
| github.com/containernetworking/plugins                                      | v1.1.1                                            |                                    | true   | true             |
| github.com/containers/buildah                                               | v1.26.2                                           |                                    | true   | true             |
| github.com/containers/common                                                | v0.48.0                                           |                                    | true   | true             |
| github.com/containers/conmon                                                | v2.0.20+incompatible                              |                                    | true   | true             |
| github.com/containers/conmon-rs                                             | v0.0.0-20220609131637-836ba11bf0aa                | v0.0.0-20220712072830-20eb3b35a5cd | true   | true             |
| github.com/containers/image/v5                                              | v5.21.1                                           |                                    | true   | true             |
| github.com/containers/libtrust                                              | v0.0.0-20200511145503-9c3a6c22cd9a                |                                    | false  | true             |
| github.com/containers/ocicrypt                                              | v1.1.5                                            |                                    | true   | true             |
| github.com/containers/podman/v4                                             | v4.1.1                                            |                                    | true   | true             |
| github.com/containers/psgo                                                  | v1.7.2                                            |                                    | false  | true             |
| github.com/containers/storage                                               | v1.41.0                                           |                                    | true   | true             |
| github.com/coredns/caddy                                                    | v1.1.0                                            | v1.1.1                             | false  | true             |
| github.com/coredns/corefile-migration                                       | v1.0.17                                           |                                    | false  | true             |
| github.com/coreos/bbolt                                                     | v1.3.2                                            | v1.3.6                             | false  | true             |
| github.com/coreos/etcd                                                      | v3.3.13+incompatible                              | v3.3.27+incompatible               | false  | true             |
| github.com/coreos/go-etcd                                                   | v2.0.0+incompatible                               |                                    | false  | true             |
| github.com/coreos/go-iptables                                               | v0.6.0                                            |                                    | false  | true             |
| github.com/coreos/go-oidc                                                   | v2.1.0+incompatible                               | v2.2.1+incompatible                | false  | true             |
| github.com/coreos/go-oidc/v3                                                | v3.1.0                                            | v3.2.0                             | false  | true             |
| github.com/coreos/go-semver                                                 | v0.3.0                                            |                                    | false  | true             |
| github.com/coreos/go-systemd                                                | v0.0.0-20190620071333-e64a0ec8b42a                | v0.0.0-20191104093116-d3cd4ed1dbcf | false  | true             |
| github.com/coreos/go-systemd/v22                                            | v22.3.2                                           |                                    | true   | true             |
| github.com/coreos/pkg                                                       | v0.0.0-20180928190104-399ea9e2e55f                | v0.0.0-20220709002704-04386ae12ed0 | false  | true             |
| github.com/coreos/stream-metadata-go                                        | v0.0.0-20210225230131-70edb9eb47b3                | v0.3.0                             | false  | true             |
| github.com/cpuguy83/go-md2man                                               | v1.0.10                                           |                                    | true   | true             |
| github.com/cpuguy83/go-md2man/v2                                            | v2.0.2                                            |                                    | false  | true             |
| github.com/creack/pty                                                       | v1.1.18                                           |                                    | true   | true             |
| github.com/cri-o/ocicni                                                     | v0.4.0                                            |                                    | true   | true             |
| github.com/cyberphone/json-canonicalization                                 | v0.0.0-20210823021906-dc406ceaf94b                | v0.0.0-20220623050100-57a0ce2678a7 | false  | true             |
| github.com/cyphar/filepath-securejoin                                       | v0.2.3                                            |                                    | true   | true             |
| github.com/d2g/dhcp4                                                        | v0.0.0-20170904100407-a1d1b6c41b1c                |                                    | false  | true             |
| github.com/d2g/dhcp4client                                                  | v1.0.0                                            |                                    | false  | true             |
| github.com/d2g/dhcp4server                                                  | v0.0.0-20181031114812-7d4a0a7f59a5                |                                    | false  | true             |
| github.com/d2g/hardwareaddr                                                 | v0.0.0-20190221164911-e7d9fbe030e4                |                                    | false  | true             |
| github.com/daixiang0/gci                                                    | v0.2.9                                            | v0.4.1                             | false  | true             |
| github.com/danieljoos/wincred                                               | v1.1.1                                            | v1.1.2                             | false  | true             |
| github.com/davecgh/go-spew                                                  | v1.1.1                                            |                                    | false  | true             |
| github.com/daviddengcn/go-colortext                                         | v1.0.0                                            |                                    | false  | true             |
| github.com/denis-tingajkin/go-header                                        | v0.4.2                                            | v0.4.3                             | false  | true             |
| github.com/denisenkom/go-mssqldb                                            | v0.11.0                                           | v0.12.2                            | false  | true             |
| github.com/denverdino/aliyungo                                              | v0.0.0-20190125010748-a747050bb1ba                | v0.0.0-20220610083100-ab5f747cb559 | false  | true             |
| github.com/devigned/tab                                                     | v0.1.1                                            |                                    | false  | true             |
| github.com/dgraph-io/badger/v3                                              | v3.2103.2                                         |                                    | false  | true             |
| github.com/dgraph-io/ristretto                                              | v0.1.0                                            |                                    | false  | true             |
| github.com/dgrijalva/jwt-go                                                 | v3.2.0+incompatible                               |                                    | false  | true             |
| github.com/dgryski/go-farm                                                  | v0.0.0-20200201041132-a6ae2369ad13                |                                    | false  | true             |
| github.com/dgryski/go-gk                                                    | v0.0.0-20200319235926-a69029f61654                |                                    | false  | true             |
| github.com/dgryski/go-rendezvous                                            | v0.0.0-20200823014737-9f7001d12a5f                |                                    | false  | true             |
| github.com/dgryski/go-sip13                                                 | v0.0.0-20181026042036-e10d5fee7954                | v0.0.0-20200911182023-62edffca9245 | false  | true             |
| github.com/digitalocean/go-libvirt                                          | v0.0.0-20201209184759-e2a69bcd5bd1                | v0.0.0-20220704152101-6d260d9c16a3 | false  | true             |
| github.com/digitalocean/go-qemu                                             | v0.0.0-20210326154740-ac9e0b687001                |                                    | false  | true             |
| github.com/dimchansky/utfbom                                                | v1.1.1                                            |                                    | false  | true             |
| github.com/disiqueira/gotree/v3                                             | v3.0.2                                            |                                    | false  | true             |
| github.com/dnaeon/go-vcr                                                    | v1.0.1                                            | v1.2.0                             | false  | true             |
| github.com/docker/cli                                                       | v20.10.16+incompatible                            | v20.10.17+incompatible             | false  | true             |
| github.com/docker/distribution                                              | v2.8.1+incompatible                               |                                    | true   | true             |
| github.com/docker/docker                                                    | v20.10.16+incompatible                            | v20.10.17+incompatible             | false  | true             |
| github.com/docker/docker-credential-helpers                                 | v0.6.4                                            |                                    | false  | true             |
| github.com/docker/go-connections                                            | v0.4.1-0.20210727194412-58542c764a11              |                                    | false  | true             |
| github.com/docker/go-events                                                 | v0.0.0-20190806004212-e31b211e4f1c                |                                    | false  | true             |
| github.com/docker/go-metrics                                                | v0.0.1                                            |                                    | false  | true             |
| github.com/docker/go-plugins-helpers                                        | v0.0.0-20211224144127-6eecb7beb651                |                                    | false  | true             |
| github.com/docker/go-units                                                  | v0.4.0                                            |                                    | true   | true             |
| github.com/docker/libnetwork                                                | v0.8.0-dev.2.0.20190625141545-5a177b73e316        |                                    | false  | true             |
| github.com/docker/libtrust                                                  | v0.0.0-20160708172513-aabc10ec26b7                |                                    | false  | true             |
| github.com/docopt/docopt-go                                                 | v0.0.0-20180111231733-ee0de3bc6815                |                                    | false  | true             |
| github.com/dsnet/compress                                                   | v0.0.2-0.20210315054119-f66993602bf5              |                                    | false  | true             |
| github.com/dtylman/scp                                                      | v0.0.0-20181017070807-f3000a34aef4                | v0.0.0-20220201143837-3fe2c6f3fee9 | false  | true             |
| github.com/dustin/go-humanize                                               | v1.0.0                                            |                                    | false  | true             |
| github.com/dvyukov/go-fuzz                                                  | v0.0.0-20210914135545-4980593459a1                | v0.0.0-20220220162807-a217d9bdbece | false  | true             |
| github.com/eapache/go-resiliency                                            | v1.1.0                                            | v1.3.0                             | false  | true             |
| github.com/eapache/go-xerial-snappy                                         | v0.0.0-20180814174437-776d5712da21                |                                    | false  | true             |
| github.com/eapache/queue                                                    | v1.1.0                                            |                                    | false  | true             |
| github.com/edsrzf/mmap-go                                                   | v1.0.0                                            | v1.1.0                             | false  | true             |
| github.com/eggsampler/acme/v3                                               | v3.2.1                                            | v3.3.0                             | false  | true             |
| github.com/elazarl/goproxy                                                  | v0.0.0-20180725130230-947c36da3153                | v0.0.0-20220529153421-8ea89ba92021 | false  | true             |
| github.com/emicklei/go-restful                                              | v2.15.0+incompatible                              | v2.16.0+incompatible               | true   | true             |
| github.com/emicklei/go-restful/v3                                           | v3.8.0                                            |                                    | false  | true             |
| github.com/emicklei/proto                                                   | v1.6.15                                           | v1.11.0                            | false  | true             |
| github.com/emirpasic/gods                                                   | v1.12.0                                           | v1.18.1                            | false  | true             |
| github.com/envoyproxy/go-control-plane                                      | v0.10.2-0.20220325020618-49ff273808a1             | v0.10.3                            | false  | true             |
| github.com/envoyproxy/protoc-gen-validate                                   | v0.6.2                                            | v0.6.7                             | false  | true             |
| github.com/esimonov/ifshort                                                 | v1.0.3                                            | v1.0.4                             | false  | true             |
| github.com/ettle/strcase                                                    | v0.1.1                                            |                                    | false  | true             |
| github.com/euank/go-kmsg-parser                                             | v2.0.0+incompatible                               |                                    | false  | true             |
| github.com/evanphx/json-patch                                               | v4.12.0+incompatible                              | v5.6.0+incompatible                | false  | false            |
| github.com/evanphx/json-patch/v5                                            | v5.6.0                                            |                                    | false  | true             |
| github.com/exponent-io/jsonpath                                             | v0.0.0-20151013193312-d6023ce2651d                | v0.0.0-20210407135951-1de76d718b3f | false  | true             |
| github.com/facebookgo/clock                                                 | v0.0.0-20150410010913-600d898af40a                |                                    | false  | true             |
| github.com/facebookgo/limitgroup                                            | v0.0.0-20150612190941-6abd8d71ec01                |                                    | false  | true             |
| github.com/facebookgo/muster                                                | v0.0.0-20150708232844-fd3d7953fd52                |                                    | false  | true             |
| github.com/fanliao/go-promise                                               | v0.0.0-20141029170127-1890db352a72                |                                    | false  | true             |
| github.com/fatih/camelcase                                                  | v1.0.0                                            |                                    | false  | true             |
| github.com/fatih/color                                                      | v1.13.0                                           |                                    | false  | true             |
| github.com/fatih/structs                                                    | v1.1.0                                            |                                    | false  | true             |
| github.com/fatih/structtag                                                  | v1.2.0                                            |                                    | false  | true             |
| github.com/felixge/httpsnoop                                                | v1.0.1                                            | v1.0.3                             | false  | true             |
| github.com/flynn/go-docopt                                                  | v0.0.0-20140912013429-f6dd2ebbb31e                |                                    | false  | true             |
| github.com/flynn/go-shlex                                                   | v0.0.0-20150515145356-3f9db97f8568                |                                    | false  | true             |
| github.com/fogleman/gg                                                      | v1.2.1-0.20190220221249-0403632d5b90              | v1.3.0                             | false  | true             |
| github.com/form3tech-oss/jwt-go                                             | v3.2.5+incompatible                               |                                    | false  | true             |
| github.com/fortytw2/leaktest                                                | v1.3.0                                            |                                    | false  | true             |
| github.com/foxcpp/go-mockdns                                                | v0.0.0-20210729171921-fb145fc6f897                | v1.0.0                             | false  | true             |
| github.com/franela/goblin                                                   | v0.0.0-20200105215937-c9ffbefa60db                | v0.0.0-20211003143422-0a4f594942bf | false  | true             |
| github.com/franela/goreq                                                    | v0.0.0-20171204163338-bcd34c9993f8                |                                    | false  | true             |
| github.com/frankban/quicktest                                               | v1.14.3                                           |                                    | false  | true             |
| github.com/fsnotify/fsnotify                                                | v1.5.4                                            |                                    | true   | true             |
| github.com/fsouza/go-dockerclient                                           | v1.7.11                                           | v1.8.1                             | false  | true             |
| github.com/fullsailor/pkcs7                                                 | v0.0.0-20190404230743-d7302db945fa                |                                    | false  | true             |
| github.com/fullstorydev/grpcurl                                             | v1.8.2                                            | v1.8.6                             | false  | true             |
| github.com/fvbommel/sortorder                                               | v1.0.1                                            | v1.0.2                             | false  | true             |
| github.com/fzipp/gocyclo                                                    | v0.3.1                                            | v0.6.0                             | false  | true             |
| github.com/garyburd/redigo                                                  | v0.0.0-20150301180006-535138d7bcd7                | v1.6.3                             | false  | true             |
| github.com/getkin/kin-openapi                                               | v0.76.0                                           | v0.97.0                            | false  | true             |
| github.com/getsentry/raven-go                                               | v0.2.0                                            |                                    | false  | true             |
| github.com/ghodss/yaml                                                      | v1.0.0                                            |                                    | false  | true             |
| github.com/gin-contrib/sse                                                  | v0.1.0                                            |                                    | false  | true             |
| github.com/gin-gonic/gin                                                    | v1.7.3                                            | v1.8.1                             | false  | true             |
| github.com/gliderlabs/ssh                                                   | v0.2.2                                            | v0.3.4                             | false  | true             |
| github.com/globalsign/mgo                                                   | v0.0.0-20181015135952-eeefdecb41b8                |                                    | false  | true             |
| github.com/go-asn1-ber/asn1-ber                                             | v1.3.1                                            | v1.5.4                             | false  | true             |
| github.com/go-chi/chi                                                       | v4.1.2+incompatible                               |                                    | false  | true             |
| github.com/go-critic/go-critic                                              | v0.6.1                                            | v0.6.3                             | false  | true             |
| github.com/go-errors/errors                                                 | v1.0.1                                            | v1.4.2                             | false  | true             |
| github.com/go-git/gcfg                                                      | v1.5.0                                            |                                    | false  | true             |
| github.com/go-git/go-billy/v5                                               | v5.3.1                                            |                                    | false  | true             |
| github.com/go-git/go-git-fixtures/v4                                        | v4.2.1                                            | v4.3.1                             | false  | true             |
| github.com/go-git/go-git/v5                                                 | v5.4.2                                            |                                    | false  | true             |
| github.com/go-gl/glfw                                                       | v0.0.0-20190409004039-e6da0acd62b1                | v0.0.0-20220712193148-63cf1f4ef61f | false  | true             |
| github.com/go-gl/glfw/v3.3/glfw                                             | v0.0.0-20200222043503-6f7a984d4dc4                | v0.0.0-20220712193148-63cf1f4ef61f | false  | true             |
| github.com/go-gorp/gorp/v3                                                  | v3.0.2                                            |                                    | false  | true             |
| github.com/go-ini/ini                                                       | v1.25.4                                           | v1.66.6                            | false  | true             |
| github.com/go-kit/kit                                                       | v0.10.0                                           | v0.12.0                            | false  | true             |
| github.com/go-kit/log                                                       | v0.1.0                                            | v0.2.1                             | false  | true             |
| github.com/go-ldap/ldap/v3                                                  | v3.1.10                                           | v3.4.3                             | false  | true             |
| github.com/go-lintpack/lintpack                                             | v0.5.2                                            |                                    | false  | true             |
| github.com/go-logfmt/logfmt                                                 | v0.5.0                                            | v0.5.1                             | false  | true             |
| github.com/go-logr/logr                                                     | v1.2.3                                            |                                    | true   | true             |
| github.com/go-logr/stdr                                                     | v1.2.2                                            |                                    | false  | true             |
| github.com/go-logr/zapr                                                     | v1.2.3                                            |                                    | false  | true             |
| github.com/go-ole/go-ole                                                    | v1.2.6                                            |                                    | false  | true             |
| github.com/go-openapi/analysis                                              | v0.21.2                                           | v0.21.3                            | false  | true             |
| github.com/go-openapi/errors                                                | v0.20.2                                           |                                    | false  | true             |
| github.com/go-openapi/jsonpointer                                           | v0.19.5                                           |                                    | false  | true             |
| github.com/go-openapi/jsonreference                                         | v0.19.6                                           | v0.20.0                            | false  | true             |
| github.com/go-openapi/loads                                                 | v0.21.1                                           |                                    | false  | true             |
| github.com/go-openapi/runtime                                               | v0.24.1                                           |                                    | false  | true             |
| github.com/go-openapi/spec                                                  | v0.20.4                                           | v0.20.6                            | false  | true             |
| github.com/go-openapi/strfmt                                                | v0.21.2                                           |                                    | false  | true             |
| github.com/go-openapi/swag                                                  | v0.21.1                                           |                                    | false  | true             |
| github.com/go-openapi/validate                                              | v0.21.0                                           | v0.22.0                            | false  | true             |
| github.com/go-ozzo/ozzo-validation                                          | v3.5.0+incompatible                               | v3.6.0+incompatible                | false  | true             |
| github.com/go-piv/piv-go                                                    | v1.9.0                                            |                                    | false  | true             |
| github.com/go-playground/assert/v2                                          | v2.0.1                                            |                                    | false  | true             |
| github.com/go-playground/locales                                            | v0.14.0                                           |                                    | false  | true             |
| github.com/go-playground/universal-translator                               | v0.18.0                                           |                                    | false  | true             |
| github.com/go-playground/validator/v10                                      | v10.10.0                                          | v10.11.0                           | false  | true             |
| github.com/go-redis/redis                                                   | v6.15.9+incompatible                              |                                    | false  | true             |
| github.com/go-redis/redis/v8                                                | v8.11.4                                           | v8.11.5                            | false  | true             |
| github.com/go-rod/rod                                                       | v0.106.1                                          | v0.108.1                           | false  | true             |
| github.com/go-sql-driver/mysql                                              | v1.6.0                                            |                                    | false  | true             |
| github.com/go-stack/stack                                                   | v1.8.1                                            |                                    | false  | true             |
| github.com/go-task/slim-sprig                                               | v0.0.0-20210107165309-348f09dbbbc0                |                                    | false  | true             |
| github.com/go-test/deep                                                     | v1.0.8                                            |                                    | false  | true             |
| github.com/go-toolsmith/astcast                                             | v1.0.0                                            |                                    | false  | true             |
| github.com/go-toolsmith/astcopy                                             | v1.0.0                                            | v1.0.1                             | false  | true             |
| github.com/go-toolsmith/astequal                                            | v1.0.1                                            | v1.0.2                             | false  | true             |
| github.com/go-toolsmith/astfmt                                              | v1.0.0                                            |                                    | false  | true             |
| github.com/go-toolsmith/astinfo                                             | v0.0.0-20180906194353-9809ff7efb21                | v1.0.0                             | false  | true             |
| github.com/go-toolsmith/astp                                                | v1.0.0                                            |                                    | false  | true             |
| github.com/go-toolsmith/pkgload                                             | v1.0.0                                            | v1.0.1                             | false  | true             |
| github.com/go-toolsmith/strparse                                            | v1.0.0                                            |                                    | false  | true             |
| github.com/go-toolsmith/typep                                               | v1.0.2                                            |                                    | false  | true             |
| github.com/go-xmlfmt/xmlfmt                                                 | v0.0.0-20191208150333-d5b6f63a941b                | v0.0.0-20220206211657-0a94163c4677 | false  | true             |
| github.com/go-zoo/bone                                                      | v1.3.0                                            |                                    | true   | true             |
| github.com/gobuffalo/attrs                                                  | v0.0.0-20190224210810-a9411de4debd                | v1.0.2                             | false  | true             |
| github.com/gobuffalo/depgen                                                 | v0.1.0                                            | v0.2.0                             | false  | true             |
| github.com/gobuffalo/envy                                                   | v1.7.0                                            | v1.10.1                            | false  | true             |
| github.com/gobuffalo/flect                                                  | v0.2.4                                            | v0.2.5                             | false  | true             |
| github.com/gobuffalo/genny                                                  | v0.1.1                                            | v0.6.0                             | false  | true             |
| github.com/gobuffalo/gitgen                                                 | v0.0.0-20190315122116-cc086187d211                |                                    | false  | true             |
| github.com/gobuffalo/gogen                                                  | v0.1.1                                            | v0.2.0                             | false  | true             |
| github.com/gobuffalo/logger                                                 | v0.0.0-20190315122211-86e12af44bc2                | v1.0.6                             | false  | true             |
| github.com/gobuffalo/mapi                                                   | v1.0.2                                            | v1.2.1                             | false  | true             |
| github.com/gobuffalo/packd                                                  | v0.1.0                                            | v1.0.1                             | false  | true             |
| github.com/gobuffalo/packr/v2                                               | v2.2.0                                            | v2.8.3                             | false  | true             |
| github.com/gobuffalo/syncx                                                  | v0.0.0-20190224160051-33c29581e754                | v0.1.0                             | false  | true             |
| github.com/gobwas/glob                                                      | v0.2.3                                            |                                    | false  | true             |
| github.com/gobwas/httphead                                                  | v0.0.0-20180130184737-2c6c146eadee                | v0.1.0                             | false  | true             |
| github.com/gobwas/pool                                                      | v0.2.0                                            | v0.2.1                             | false  | true             |
| github.com/gobwas/ws                                                        | v1.0.2                                            | v1.1.0                             | false  | true             |
| github.com/godbus/dbus                                                      | v4.1.0+incompatible                               |                                    | false  | true             |
| github.com/godbus/dbus/v5                                                   | v5.1.0                                            |                                    | true   | true             |
| github.com/gofrs/flock                                                      | v0.8.1                                            |                                    | false  | true             |
| github.com/gofrs/uuid                                                       | v4.0.0+incompatible                               | v4.2.0+incompatible                | false  | true             |
| github.com/gogo/googleapis                                                  | v1.4.0                                            | v1.4.1                             | false  | true             |
| github.com/gogo/protobuf                                                    | v1.3.2                                            |                                    | true   | true             |
| github.com/golang-jwt/jwt                                                   | v3.2.2+incompatible                               |                                    | false  | true             |
| github.com/golang-jwt/jwt/v4                                                | v4.3.0                                            | v4.4.2                             | false  | true             |
| github.com/golang-sql/civil                                                 | v0.0.0-20190719163853-cb61b32ac6fe                | v0.0.0-20220223132316-b832511892a9 | false  | true             |
| github.com/golang/freetype                                                  | v0.0.0-20170609003504-e2365dfdc4a0                |                                    | false  | true             |
| github.com/golang/glog                                                      | v1.0.0                                            |                                    | false  | true             |
| github.com/golang/groupcache                                                | v0.0.0-20210331224755-41bb18bfe9da                |                                    | false  | true             |
| github.com/golang/mock                                                      | v1.6.0                                            |                                    | true   | true             |
| github.com/golang/protobuf                                                  | v1.5.2                                            |                                    | false  | true             |
| github.com/golang/snappy                                                    | v0.0.4                                            |                                    | false  | true             |
| github.com/golangci/check                                                   | v0.0.0-20180506172741-cfe4005ccda2                |                                    | false  | true             |
| github.com/golangci/dupl                                                    | v0.0.0-20180902072040-3e9179ac440a                |                                    | false  | true             |
| github.com/golangci/errcheck                                                | v0.0.0-20181223084120-ef45e06d44b6                |                                    | false  | true             |
| github.com/golangci/go-misc                                                 | v0.0.0-20180628070357-927a3d87b613                | v0.0.0-20220329215616-d24fe342adfe | false  | true             |
| github.com/golangci/goconst                                                 | v0.0.0-20180610141641-041c5f2b40f3                |                                    | false  | true             |
| github.com/golangci/gocyclo                                                 | v0.0.0-20180528134321-2becd97e67ee                | v0.0.0-20180528144436-0a533e8fa43d | false  | true             |
| github.com/golangci/gofmt                                                   | v0.0.0-20190930125516-244bba706f1a                |                                    | false  | true             |
| github.com/golangci/golangci-lint                                           | v1.43.0                                           | v1.46.2                            | false  | true             |
| github.com/golangci/ineffassign                                             | v0.0.0-20190609212857-42439a7714cc                |                                    | false  | true             |
| github.com/golangci/lint-1                                                  | v0.0.0-20191013205115-297bf364a8e0                |                                    | false  | true             |
| github.com/golangci/maligned                                                | v0.0.0-20180506175553-b1d89398deca                |                                    | false  | true             |
| github.com/golangci/misspell                                                | v0.3.5                                            |                                    | false  | true             |
| github.com/golangci/prealloc                                                | v0.0.0-20180630174525-215b22d4de21                |                                    | false  | true             |
| github.com/golangci/revgrep                                                 | v0.0.0-20210930125155-c22e5001d4f2                |                                    | false  | true             |
| github.com/golangci/unconvert                                               | v0.0.0-20180507085042-28b1c447d1f4                |                                    | false  | true             |
| github.com/golangplus/bytes                                                 | v1.0.0                                            |                                    | false  | true             |
| github.com/golangplus/fmt                                                   | v1.0.0                                            |                                    | false  | true             |
| github.com/golangplus/testing                                               | v1.0.0                                            |                                    | false  | true             |
| github.com/gomarkdown/markdown                                              | v0.0.0-20210514010506-3b9f47219fe7                | v0.0.0-20220627144906-e9a81102ebeb | false  | true             |
| github.com/google/btree                                                     | v1.0.1                                            | v1.1.2                             | false  | true             |
| github.com/google/cadvisor                                                  | v0.44.1                                           |                                    | false  | true             |
| github.com/google/cel-go                                                    | v0.11.2                                           | v0.12.1                            | false  | true             |
| github.com/google/certificate-transparency-go                               | v1.1.2                                            | v1.1.3                             | false  | true             |
| github.com/google/flatbuffers                                               | v1.12.1                                           | v2.0.6+incompatible                | false  | true             |
| github.com/google/gnostic                                                   | v0.5.7-v3refs                                     | v0.6.9                             | false  | true             |
| github.com/google/go-cmp                                                    | v0.5.8                                            |                                    | false  | true             |
| github.com/google/go-containerregistry                                      | v0.10.0                                           |                                    | false  | true             |
| github.com/google/go-containerregistry/pkg/authn/k8schain                   | v0.0.0-20220413173345-f1b065c6cb3d                | v0.0.0-20220712174516-ddd39fb9c385 | false  | true             |
| github.com/google/go-containerregistry/pkg/authn/kubernetes                 | v0.0.0-20220301182634-bfe2ffc6b6bd                | v0.0.0-20220712174516-ddd39fb9c385 | false  | true             |
| github.com/google/go-github/v27                                             | v27.0.6                                           |                                    | false  | true             |
| github.com/google/go-github/v28                                             | v28.1.1                                           |                                    | false  | true             |
| github.com/google/go-github/v33                                             | v33.0.0                                           |                                    | false  | true             |
| github.com/google/go-github/v39                                             | v39.2.0                                           |                                    | false  | true             |
| github.com/google/go-github/v42                                             | v42.0.0                                           |                                    | false  | true             |
| github.com/google/go-github/v45                                             | v45.1.0                                           | v45.2.0                            | false  | true             |
| github.com/google/go-intervals                                              | v0.0.2                                            |                                    | false  | true             |
| github.com/google/go-licenses                                               | v0.0.0-20210329231322-ce1d9163b77d                | v1.2.1                             | false  | true             |
| github.com/google/go-querystring                                            | v1.1.0                                            |                                    | false  | true             |
| github.com/google/go-replayers/grpcreplay                                   | v1.1.0                                            |                                    | false  | true             |
| github.com/google/go-replayers/httpreplay                                   | v1.0.0                                            | v1.1.1                             | false  | true             |
| github.com/google/gofuzz                                                    | v1.2.0                                            |                                    | false  | true             |
| github.com/google/licenseclassifier                                         | v0.0.0-20210325184830-bb04aff29e72                | v0.0.0-20220326190949-7c62d6fe8d3a | false  | true             |
| github.com/google/licenseclassifier/v2                                      | v2.0.0-alpha.1                                    | v2.0.0-pre5                        | false  | true             |
| github.com/google/mako                                                      | v0.0.0-20190821191249-122f8dcef9e3                | v0.2.0                             | false  | true             |
| github.com/google/martian                                                   | v2.1.1-0.20190517191504-25dcb96d9e51+incompatible |                                    | false  | true             |
| github.com/google/martian/v3                                                | v3.2.1                                            | v3.3.2                             | false  | true             |
| github.com/google/pprof                                                     | v0.0.0-20210720184732-4bb14d4b1be1                | v0.0.0-20220608213341-c488b8fa1db3 | false  | true             |
| github.com/google/renameio                                                  | v1.0.1                                            |                                    | true   | true             |
| github.com/google/rpmpack                                                   | v0.0.0-20210518075352-dc539ef4f2ea                | v0.0.0-20220411070212-51a1004ef6cb | false  | true             |
| github.com/google/shlex                                                     | v0.0.0-20191202100458-e7afc7fbc510                |                                    | false  | true             |
| github.com/google/subcommands                                               | v1.0.1                                            | v1.2.0                             | false  | true             |
| github.com/google/trillian                                                  | v1.4.0                                            | v1.4.1                             | false  | true             |
| github.com/google/uuid                                                      | v1.3.0                                            |                                    | true   | true             |
| github.com/google/wire                                                      | v0.5.0                                            |                                    | false  | true             |
| github.com/googleapis/gax-go                                                | v2.0.2+incompatible                               |                                    | false  | true             |
| github.com/googleapis/gax-go/v2                                             | v2.4.0                                            |                                    | false  | true             |
| github.com/googleapis/gnostic                                               | v0.5.5                                            | v0.6.9                             | false  | true             |
| github.com/googleapis/go-type-adapters                                      | v1.0.0                                            |                                    | false  | true             |
| github.com/googleapis/google-cloud-go-testing                               | v0.0.0-20200911160855-bcd43fbb19e8                | v0.0.0-20210719221736-1c9a4c676720 | false  | true             |
| github.com/gookit/color                                                     | v1.4.2                                            | v1.5.1                             | false  | true             |
| github.com/gophercloud/gophercloud                                          | v0.1.0                                            | v0.25.0                            | false  | true             |
| github.com/gopherjs/gopherjs                                                | v0.0.0-20200217142428-fce0ec30dd00                | v1.17.2                            | false  | true             |
| github.com/gordonklaus/ineffassign                                          | v0.0.0-20210225214923-2e10b2664254                | v0.0.0-20210914165742-4cc7213b9bc8 | false  | true             |
| github.com/goreleaser/goreleaser                                            | v0.134.0                                          | v1.10.2                            | false  | true             |
| github.com/goreleaser/nfpm                                                  | v1.2.1                                            | v1.10.3                            | false  | true             |
| github.com/gorhill/cronexpr                                                 | v0.0.0-20180427100037-88b0669f7d75                |                                    | false  | true             |
| github.com/gorilla/context                                                  | v1.1.1                                            |                                    | false  | true             |
| github.com/gorilla/handlers                                                 | v1.5.1                                            |                                    | false  | true             |
| github.com/gorilla/mux                                                      | v1.8.0                                            |                                    | false  | true             |
| github.com/gorilla/schema                                                   | v1.2.0                                            |                                    | false  | true             |
| github.com/gorilla/websocket                                                | v1.4.2                                            | v1.5.0                             | false  | true             |
| github.com/gostaticanalysis/analysisutil                                    | v0.7.1                                            |                                    | false  | true             |
| github.com/gostaticanalysis/comment                                         | v1.4.2                                            |                                    | false  | true             |
| github.com/gostaticanalysis/forcetypeassert                                 | v0.0.0-20200621232751-01d4955beaa5                | v0.1.0                             | false  | true             |
| github.com/gostaticanalysis/nilerr                                          | v0.1.1                                            |                                    | false  | true             |
| github.com/gostaticanalysis/testutil                                        | v0.4.0                                            |                                    | false  | true             |
| github.com/gregjones/httpcache                                              | v0.0.0-20190611155906-901d90724c79                |                                    | false  | true             |
| github.com/grpc-ecosystem/go-grpc-middleware                                | v1.3.0                                            |                                    | true   | true             |
| github.com/grpc-ecosystem/go-grpc-prometheus                                | v1.2.0                                            |                                    | false  | true             |
| github.com/grpc-ecosystem/grpc-gateway                                      | v1.16.0                                           |                                    | false  | true             |
| github.com/grpc-ecosystem/grpc-gateway/v2                                   | v2.7.0                                            | v2.10.3                            | false  | true             |
| github.com/hanwen/go-fuse                                                   | v1.0.0                                            |                                    | false  | true             |
| github.com/hanwen/go-fuse/v2                                                | v2.1.0                                            |                                    | false  | true             |
| github.com/hashicorp/consul/api                                             | v1.12.0                                           | v1.13.0                            | false  | true             |
| github.com/hashicorp/consul/sdk                                             | v0.8.0                                            | v0.9.0                             | false  | true             |
| github.com/hashicorp/errwrap                                                | v1.1.0                                            |                                    | false  | true             |
| github.com/hashicorp/go-cleanhttp                                           | v0.5.2                                            |                                    | false  | true             |
| github.com/hashicorp/go-hclog                                               | v1.2.0                                            | v1.2.1                             | false  | true             |
| github.com/hashicorp/go-immutable-radix                                     | v1.3.1                                            |                                    | false  | true             |
| github.com/hashicorp/go-kms-wrapping/entropy                                | v0.1.0                                            |                                    | false  | true             |
| github.com/hashicorp/go-msgpack                                             | v0.5.3                                            | v1.1.5                             | false  | true             |
| github.com/hashicorp/go-multierror                                          | v1.1.1                                            |                                    | false  | true             |
| github.com/hashicorp/go-plugin                                              | v1.4.4                                            |                                    | false  | true             |
| github.com/hashicorp/go-retryablehttp                                       | v0.7.1                                            |                                    | false  | true             |
| github.com/hashicorp/go-rootcerts                                           | v1.0.2                                            |                                    | false  | true             |
| github.com/hashicorp/go-secure-stdlib/base62                                | v0.1.1                                            | v0.1.2                             | false  | true             |
| github.com/hashicorp/go-secure-stdlib/mlock                                 | v0.1.2                                            |                                    | false  | true             |
| github.com/hashicorp/go-secure-stdlib/parseutil                             | v0.1.5                                            | v0.1.6                             | false  | true             |
| github.com/hashicorp/go-secure-stdlib/password                              | v0.1.1                                            |                                    | false  | true             |
| github.com/hashicorp/go-secure-stdlib/strutil                               | v0.1.2                                            |                                    | false  | true             |
| github.com/hashicorp/go-secure-stdlib/tlsutil                               | v0.1.1                                            |                                    | false  | true             |
| github.com/hashicorp/go-sockaddr                                            | v1.0.2                                            |                                    | false  | true             |
| github.com/hashicorp/go-syslog                                              | v1.0.0                                            |                                    | false  | true             |
| github.com/hashicorp/go-uuid                                                | v1.0.3                                            |                                    | false  | true             |
| github.com/hashicorp/go-version                                             | v1.5.0                                            | v1.6.0                             | false  | true             |
| github.com/hashicorp/go.net                                                 | v0.0.1                                            |                                    | false  | true             |
| github.com/hashicorp/golang-lru                                             | v0.5.4                                            |                                    | false  | true             |
| github.com/hashicorp/hcl                                                    | v1.0.0                                            |                                    | false  | true             |
| github.com/hashicorp/logutils                                               | v1.0.0                                            |                                    | false  | true             |
| github.com/hashicorp/mdns                                                   | v1.0.4                                            | v1.0.5                             | false  | true             |
| github.com/hashicorp/memberlist                                             | v0.3.0                                            | v0.3.1                             | false  | true             |
| github.com/hashicorp/serf                                                   | v0.9.7                                            | v0.9.8                             | false  | true             |
| github.com/hashicorp/vault/api                                              | v1.5.0                                            | v1.7.2                             | false  | true             |
| github.com/hashicorp/vault/sdk                                              | v0.5.0                                            | v0.5.2                             | false  | true             |
| github.com/hashicorp/yamux                                                  | v0.0.0-20211028200310-0bc27b27de87                |                                    | false  | true             |
| github.com/heketi/heketi                                                    | v10.3.0+incompatible                              |                                    | false  | true             |
| github.com/heketi/tests                                                     | v0.0.0-20151005000721-f3775cbcefd6                |                                    | false  | true             |
| github.com/honeycombio/beeline-go                                           | v1.1.1                                            | v1.9.0                             | false  | true             |
| github.com/honeycombio/libhoney-go                                          | v1.15.2                                           | v1.15.8                            | false  | true             |
| github.com/howeyc/gopass                                                    | v0.0.0-20190910152052-7cb4b85ec19c                | v0.0.0-20210920133722-c8aef6fb66ef | false  | true             |
| github.com/hpcloud/tail                                                     | v1.0.0                                            |                                    | false  | true             |
| github.com/huandu/xstrings                                                  | v1.3.2                                            |                                    | false  | true             |
| github.com/hudl/fargo                                                       | v1.3.0                                            | v1.4.0                             | false  | true             |
| github.com/hugelgupf/socketpair                                             | v0.0.0-20190730060125-05d35a94e714                |                                    | false  | true             |
| github.com/iancoleman/strcase                                               | v0.2.0                                            |                                    | false  | true             |
| github.com/ianlancetaylor/demangle                                          | v0.0.0-20200824232613-28f6c0f3b639                | v0.0.0-20220517205856-0058ec4f073c | false  | true             |
| github.com/imdario/mergo                                                    | v0.3.12                                           | v0.3.13                            | false  | true             |
| github.com/in-toto/in-toto-golang                                           | v0.3.4-0.20211211042327-af1f9fb822bf              |                                    | false  | true             |
| github.com/inconshreveable/mousetrap                                        | v1.0.0                                            |                                    | false  | true             |
| github.com/influxdata/influxdb1-client                                      | v0.0.0-20191209144304-8bf82d3c094d                | v0.0.0-20220302092344-a9ab5670611c | false  | true             |
| github.com/influxdata/tdigest                                               | v0.0.0-20180711151920-a7d76c6f093a                | v0.0.1                             | false  | true             |
| github.com/insomniacslk/dhcp                                                | v0.0.0-20220119180841-3c283ff8b7dd                | v0.0.0-20220504074936-1ca156eafb9f | false  | true             |
| github.com/intel/goresctrl                                                  | v0.2.0                                            |                                    | true   | true             |
| github.com/ishidawataru/sctp                                                | v0.0.0-20210226210310-f2269e66cdee                | v0.0.0-20210707070123-9a39160e9062 | false  | true             |
| github.com/j-keck/arping                                                    | v1.0.2                                            |                                    | false  | true             |
| github.com/jarcoal/httpmock                                                 | v1.0.5                                            | v1.2.0                             | false  | true             |
| github.com/jbenet/go-context                                                | v0.0.0-20150711004518-d14ea06fba99                |                                    | false  | true             |
| github.com/jedisct1/go-minisign                                             | v0.0.0-20211028175153-1c139d1cc84b                |                                    | false  | true             |
| github.com/jessevdk/go-flags                                                | v1.5.0                                            |                                    | false  | true             |
| github.com/jgautheron/goconst                                               | v1.5.1                                            |                                    | false  | true             |
| github.com/jhump/protoreflect                                               | v1.9.0                                            | v1.12.0                            | false  | true             |
| github.com/jingyugao/rowserrcheck                                           | v1.1.1                                            |                                    | false  | true             |
| github.com/jinzhu/copier                                                    | v0.3.5                                            |                                    | false  | true             |
| github.com/jirfag/go-printf-func-name                                       | v0.0.0-20200119135958-7558a9eaa5af                |                                    | false  | true             |
| github.com/jmespath/go-jmespath                                             | v0.4.0                                            |                                    | false  | true             |
| github.com/jmespath/go-jmespath/internal/testify                            | v1.5.1                                            |                                    | false  | true             |
| github.com/jmhodges/clock                                                   | v0.0.0-20160418191101-880ee4c33548                | v1.2.0                             | false  | true             |
| github.com/jmoiron/sqlx                                                     | v1.2.1-0.20190826204134-d7d95172beb5              | v1.3.5                             | false  | true             |
| github.com/joefitzgerald/rainbow-reporter                                   | v0.1.0                                            |                                    | false  | true             |
| github.com/joho/godotenv                                                    | v1.3.0                                            | v1.4.0                             | false  | true             |
| github.com/jonboulle/clockwork                                              | v0.2.2                                            | v0.3.0                             | false  | true             |
| github.com/josharian/intern                                                 | v1.0.0                                            |                                    | false  | true             |
| github.com/josharian/txtarfs                                                | v0.0.0-20210218200122-0702f000015a                | v0.0.0-20210615234325-77aca6df5bca | false  | true             |
| github.com/jpillora/backoff                                                 | v1.0.0                                            |                                    | false  | true             |
| github.com/jsimonetti/rtnetlink                                             | v0.0.0-20201110080708-d2c240429e6c                | v1.2.0                             | false  | true             |
| github.com/json-iterator/go                                                 | v1.1.12                                           |                                    | true   | true             |
| github.com/jstemmer/go-junit-report                                         | v0.9.1                                            | v1.0.0                             | false  | true             |
| github.com/jtolds/gls                                                       | v4.20.0+incompatible                              |                                    | false  | true             |
| github.com/juju/ratelimit                                                   | v1.0.1                                            |                                    | false  | true             |
| github.com/julienschmidt/httprouter                                         | v1.3.0                                            |                                    | false  | true             |
| github.com/julz/importas                                                    | v0.0.0-20210419104244-841f0c0fe66d                | v0.1.0                             | false  | true             |
| github.com/jung-kurt/gofpdf                                                 | v1.0.3-0.20190309125859-24315acbbda5              | v1.16.2                            | false  | true             |
| github.com/k0kubun/colorstring                                              | v0.0.0-20150214042306-9440f1994b88                |                                    | false  | true             |
| github.com/karrick/godirwalk                                                | v1.16.1                                           | v1.17.0                            | false  | true             |
| github.com/kelseyhightower/envconfig                                        | v1.4.0                                            |                                    | false  | true             |
| github.com/kevinburke/ssh_config                                            | v1.1.0                                            | v1.2.0                             | false  | true             |
| github.com/kisielk/errcheck                                                 | v1.6.0                                            | v1.6.1                             | false  | true             |
| github.com/kisielk/gotool                                                   | v1.0.0                                            |                                    | false  | true             |
| github.com/klauspost/compress                                               | v1.15.4                                           | v1.15.7                            | false  | true             |
| github.com/klauspost/cpuid                                                  | v1.2.0                                            | v1.3.1                             | false  | true             |
| github.com/klauspost/pgzip                                                  | v1.2.5                                            |                                    | false  | true             |
| github.com/konsorten/go-windows-terminal-sequences                          | v1.0.3                                            |                                    | false  | true             |
| github.com/kr/fs                                                            | v0.1.0                                            |                                    | false  | true             |
| github.com/kr/logfmt                                                        | v0.0.0-20140226030751-b84e30acd515                | v0.0.0-20210122060352-19f9bcb100e6 | false  | true             |
| github.com/kr/pretty                                                        | v0.3.0                                            |                                    | false  | true             |
| github.com/kr/pty                                                           | v1.1.8                                            |                                    | false  | true             |
| github.com/kr/text                                                          | v0.2.0                                            |                                    | false  | true             |
| github.com/kulti/thelper                                                    | v0.4.0                                            | v0.6.3                             | false  | true             |
| github.com/kunwardeep/paralleltest                                          | v1.0.3                                            | v1.0.6                             | false  | true             |
| github.com/kylelemons/godebug                                               | v1.1.0                                            |                                    | false  | true             |
| github.com/kyoh86/exportloopref                                             | v0.1.8                                            |                                    | false  | true             |
| github.com/ldez/gomoddirectives                                             | v0.2.2                                            | v0.2.3                             | false  | true             |
| github.com/ldez/tagliatelle                                                 | v0.2.0                                            | v0.3.1                             | false  | true             |
| github.com/leodido/go-urn                                                   | v1.2.1                                            |                                    | false  | true             |
| github.com/letsencrypt/boulder                                              | v0.0.0-20220331220046-b23ab962616e                | v0.0.0-20220708192244-c7014dfd29c7 | false  | true             |
| github.com/letsencrypt/challtestsrv                                         | v1.2.1                                            |                                    | false  | true             |
| github.com/letsencrypt/pkcs11key/v4                                         | v4.0.0                                            |                                    | false  | true             |
| github.com/lib/pq                                                           | v1.10.4                                           | v1.10.6                            | false  | true             |
| github.com/libopenstorage/openstorage                                       | v1.0.0                                            | v9.4.21+incompatible               | false  | true             |
| github.com/liggitt/tabwriter                                                | v0.0.0-20181228230101-89fcab3d43de                |                                    | false  | true             |
| github.com/lightstep/lightstep-tracer-common/golang/gogo                    | v0.0.0-20190605223551-bc2310a04743                | v0.0.0-20210210170715-a8dfcb80d3a7 | false  | true             |
| github.com/lightstep/lightstep-tracer-go                                    | v0.18.1                                           | v0.25.0                            | false  | true             |
| github.com/linuxkit/virtsock                                                | v0.0.0-20201010232012-f8cee7dfc7a3                | v0.0.0-20220523201153-1a23e78aa7a2 | false  | true             |
| github.com/lithammer/dedent                                                 | v1.1.0                                            |                                    | false  | true             |
| github.com/logrusorgru/aurora                                               | v0.0.0-20181002194514-a7b3b318ed4e                | v2.0.3+incompatible                | false  | true             |
| github.com/lpabon/godbc                                                     | v0.1.1                                            |                                    | false  | true             |
| github.com/lufia/plan9stats                                                 | v0.0.0-20211012122336-39d0f177ccd0                | v0.0.0-20220517141722-cf486979b281 | false  | true             |
| github.com/lyft/protoc-gen-star                                             | v0.5.3                                            | v0.6.0                             | false  | true             |
| github.com/lyft/protoc-gen-validate                                         | v0.0.13                                           | v0.6.7                             | false  | true             |
| github.com/magefile/mage                                                    | v1.13.0                                           |                                    | false  | true             |
| github.com/magiconair/properties                                            | v1.8.6                                            |                                    | false  | true             |
| github.com/mailru/easyjson                                                  | v0.7.7                                            |                                    | false  | true             |
| github.com/manifoldco/promptui                                              | v0.9.0                                            |                                    | false  | true             |
| github.com/maratori/testpackage                                             | v1.0.1                                            | v1.1.0                             | false  | true             |
| github.com/markbates/oncer                                                  | v0.0.0-20181203154359-bf2de49a0be2                | v1.0.0                             | false  | true             |
| github.com/markbates/safe                                                   | v1.0.1                                            |                                    | false  | true             |
| github.com/marstr/guid                                                      | v1.1.0                                            |                                    | false  | true             |
| github.com/matoous/godox                                                    | v0.0.0-20210227103229-6504466cf951                |                                    | false  | true             |
| github.com/matryer/is                                                       | v1.4.0                                            |                                    | false  | true             |
| github.com/mattn/go-colorable                                               | v0.1.12                                           |                                    | false  | true             |
| github.com/mattn/go-ieproxy                                                 | v0.0.1                                            | v0.0.7                             | false  | true             |
| github.com/mattn/go-isatty                                                  | v0.0.14                                           |                                    | false  | true             |
| github.com/mattn/go-runewidth                                               | v0.0.13                                           |                                    | false  | true             |
| github.com/mattn/go-shellwords                                              | v1.0.12                                           |                                    | false  | true             |
| github.com/mattn/go-sqlite3                                                 | v1.9.0                                            | v1.14.14                           | false  | true             |
| github.com/mattn/go-zglob                                                   | v0.0.1                                            | v0.0.3                             | false  | true             |
| github.com/mattn/goveralls                                                  | v0.0.2                                            | v0.0.11                            | false  | true             |
| github.com/matttproud/golang_protobuf_extensions                            | v1.0.2-0.20181231171920-c182affec369              |                                    | false  | true             |
| github.com/maxbrunsfeld/counterfeiter/v6                                    | v6.5.0                                            |                                    | false  | true             |
| github.com/mbilski/exhaustivestruct                                         | v1.2.0                                            |                                    | false  | true             |
| github.com/mdlayher/ethernet                                                | v0.0.0-20190606142754-0394541c37b7                | v0.0.0-20220221185849-529eae5b6118 | false  | true             |
| github.com/mdlayher/netlink                                                 | v1.1.1                                            | v1.6.0                             | false  | true             |
| github.com/mdlayher/raw                                                     | v0.0.0-20191009151244-50f2db8cc065                | v0.1.0                             | false  | true             |
| github.com/mediocregopher/radix/v4                                          | v4.0.0                                            | v4.1.0                             | false  | true             |
| github.com/mgechev/dots                                                     | v0.0.0-20210922191527-e955255bf517                |                                    | false  | true             |
| github.com/mgechev/revive                                                   | v1.1.2                                            | v1.2.1                             | false  | true             |
| github.com/mgutz/ansi                                                       | v0.0.0-20170206155736-9520e82c474b                | v0.0.0-20200706080929-d51e80ef957d | false  | true             |
| github.com/mholt/archiver/v3                                                | v3.5.1                                            |                                    | false  | true             |
| github.com/miekg/dns                                                        | v1.1.45                                           | v1.1.50                            | false  | true             |
| github.com/miekg/pkcs11                                                     | v1.1.1                                            |                                    | false  | true             |
| github.com/mindprince/gonvml                                                | v0.0.0-20190828220739-9ebdce4bb989                | v0.0.0-20211002210717-ac0b66419a41 | false  | true             |
| github.com/mistifyio/go-zfs                                                 | v2.1.2-0.20190413222219-f784269be439+incompatible |                                    | false  | true             |
| github.com/mitchellh/cli                                                    | v1.1.0                                            | v1.1.4                             | false  | true             |
| github.com/mitchellh/copystructure                                          | v1.2.0                                            |                                    | false  | true             |
| github.com/mitchellh/go-homedir                                             | v1.1.0                                            |                                    | false  | true             |
| github.com/mitchellh/go-ps                                                  | v1.0.0                                            |                                    | false  | true             |
| github.com/mitchellh/go-testing-interface                                   | v1.14.1                                           |                                    | false  | true             |
| github.com/mitchellh/go-wordwrap                                            | v1.0.0                                            | v1.0.1                             | false  | true             |
| github.com/mitchellh/gox                                                    | v0.4.0                                            | v1.0.1                             | false  | true             |
| github.com/mitchellh/iochan                                                 | v1.0.0                                            |                                    | false  | true             |
| github.com/mitchellh/mapstructure                                           | v1.5.0                                            |                                    | false  | true             |
| github.com/mitchellh/osext                                                  | v0.0.0-20151018003038-5e2d6d41470f                |                                    | false  | true             |
| github.com/mitchellh/reflectwalk                                            | v1.0.2                                            |                                    | false  | true             |
| github.com/mmarkdown/mmark                                                  | v2.0.40+incompatible                              |                                    | false  | true             |
| github.com/mndrix/tap-go                                                    | v0.0.0-20171203230836-629fa407e90b                |                                    | false  | true             |
| github.com/moby/ipvs                                                        | v1.0.1                                            | v1.0.2                             | false  | true             |
| github.com/moby/locker                                                      | v1.0.1                                            |                                    | false  | true             |
| github.com/moby/spdystream                                                  | v0.2.0                                            |                                    | false  | true             |
| github.com/moby/sys/mount                                                   | v0.2.0                                            | v0.3.3                             | false  | true             |
| github.com/moby/sys/mountinfo                                               | v0.6.1                                            | v0.6.2                             | false  | true             |
| github.com/moby/sys/signal                                                  | v0.6.0                                            | v0.7.0                             | false  | true             |
| github.com/moby/sys/symlink                                                 | v0.2.0                                            |                                    | false  | true             |
| github.com/moby/term                                                        | v0.0.0-20210619224110-3f7ff695adc6                |                                    | false  | true             |
| github.com/moby/vpnkit                                                      | v0.5.0                                            |                                    | false  | true             |
| github.com/modern-go/concurrent                                             | v0.0.0-20180306012644-bacd9c7ef1dd                |                                    | false  | true             |
| github.com/modern-go/reflect2                                               | v1.0.2                                            |                                    | false  | true             |
| github.com/mohae/deepcopy                                                   | v0.0.0-20170929034955-c48cc78d4826                |                                    | false  | true             |
| github.com/monochromegane/go-gitignore                                      | v0.0.0-20200626010858-205db1a8cc00                |                                    | false  | true             |
| github.com/montanaflynn/stats                                               | v0.0.0-20171201202039-1bf9dbcd8cbe                | v0.6.6                             | false  | true             |
| github.com/moricho/tparallel                                                | v0.2.1                                            |                                    | false  | true             |
| github.com/morikuni/aec                                                     | v1.0.0                                            |                                    | false  | true             |
| github.com/mozilla/scribe                                                   | v0.0.0-20180711195314-fb71baf557c1                | v0.0.0-20220110210141-3fd4271eb395 | false  | true             |
| github.com/mozilla/tls-observatory                                          | v0.0.0-20210609171429-7bc42856d2e5                |                                    | false  | true             |
| github.com/mpvl/unique                                                      | v0.0.0-20150818121801-cbe035fff7de                |                                    | false  | true             |
| github.com/mrunalp/fileutils                                                | v0.5.0                                            |                                    | false  | true             |
| github.com/munnerz/goautoneg                                                | v0.0.0-20191010083416-a7dc8b61c822                |                                    | false  | true             |
| github.com/mvdan/xurls                                                      | v1.1.0                                            |                                    | false  | true             |
| github.com/mwitkow/go-conntrack                                             | v0.0.0-20190716064945-2f068394615f                |                                    | false  | true             |
| github.com/mwitkow/go-proto-validators                                      | v0.2.0                                            | v0.3.2                             | false  | true             |
| github.com/mxk/go-flowrate                                                  | v0.0.0-20140419014527-cca7078d478f                |                                    | false  | true             |
| github.com/nakabonne/nestif                                                 | v0.3.1                                            |                                    | false  | true             |
| github.com/nats-io/jwt                                                      | v0.3.2                                            | v1.2.2                             | false  | true             |
| github.com/nats-io/nats-server/v2                                           | v2.1.2                                            | v2.8.4                             | false  | true             |
| github.com/nats-io/nats.go                                                  | v1.9.1                                            | v1.16.0                            | false  | true             |
| github.com/nats-io/nkeys                                                    | v0.1.3                                            | v0.3.0                             | false  | true             |
| github.com/nats-io/nuid                                                     | v1.0.1                                            |                                    | false  | true             |
| github.com/nbutton23/zxcvbn-go                                              | v0.0.0-20210217022336-fa2cb2858354                |                                    | false  | true             |
| github.com/ncw/swift                                                        | v1.0.47                                           | v1.0.53                            | false  | true             |
| github.com/networkplumbing/go-nft                                           | v0.2.0                                            |                                    | false  | true             |
| github.com/niemeyer/pretty                                                  | v0.0.0-20200227124842-a10e7caefd8e                |                                    | false  | true             |
| github.com/nishanths/exhaustive                                             | v0.2.3                                            | v0.8.1                             | false  | true             |
| github.com/nishanths/predeclared                                            | v0.2.1                                            | v0.2.2                             | false  | true             |
| github.com/nozzle/throttler                                                 | v0.0.0-20180817012639-2ea982251481                |                                    | false  | true             |
| github.com/nwaples/rardecode                                                | v1.1.0                                            | v1.1.3                             | false  | true             |
| github.com/nxadm/tail                                                       | v1.4.8                                            |                                    | false  | true             |
| github.com/oklog/oklog                                                      | v0.3.2                                            |                                    | false  | true             |
| github.com/oklog/run                                                        | v1.1.0                                            |                                    | false  | true             |
| github.com/oklog/ulid                                                       | v1.3.1                                            |                                    | false  | true             |
| github.com/olekukonko/tablewriter                                           | v0.0.5                                            |                                    | false  | true             |
| github.com/onsi/ginkgo                                                      | v1.16.5                                           |                                    | false  | true             |
| github.com/onsi/ginkgo/v2                                                   | v2.1.4                                            |                                    | true   | true             |
| github.com/onsi/gomega                                                      | v1.19.0                                           |                                    | true   | true             |
| github.com/op/go-logging                                                    | v0.0.0-20160315200505-970db520ece7                |                                    | false  | true             |
| github.com/open-policy-agent/opa                                            | v0.35.0                                           | v0.42.1                            | false  | true             |
| github.com/opencontainers/go-digest                                         | v1.0.0                                            |                                    | true   | true             |
| github.com/opencontainers/image-spec                                        | v1.0.3-0.20220114050600-8b9d41f48198              |                                    | true   | true             |
| github.com/opencontainers/runc                                              | v1.1.3                                            |                                    | true   | true             |
| github.com/opencontainers/runtime-spec                                      | v1.0.3-0.20211214071223-8958f93039ab              |                                    | true   | true             |
| github.com/opencontainers/runtime-tools                                     | v0.9.1-0.20220110225228-7e2d60f1e41f              |                                    | true   | true             |
| github.com/opencontainers/selinux                                           | v1.10.1                                           |                                    | true   | true             |
| github.com/openshift/imagebuilder                                           | v1.2.4-0.20220502172744-009dbc6cb805              |                                    | false  | true             |
| github.com/opentracing-contrib/go-observer                                  | v0.0.0-20170622124052-a52f23424492                |                                    | false  | true             |
| github.com/opentracing/basictracer-go                                       | v1.0.0                                            | v1.1.0                             | false  | true             |
| github.com/opentracing/opentracing-go                                       | v1.2.0                                            |                                    | false  | true             |
| github.com/openzipkin-contrib/zipkin-go-opentracing                         | v0.4.5                                            |                                    | false  | true             |
| github.com/openzipkin/zipkin-go                                             | v0.3.0                                            | v0.4.0                             | false  | true             |
| github.com/ostreedev/ostree-go                                              | v0.0.0-20210805093236-719684c64e4f                |                                    | false  | true             |
| github.com/otiai10/copy                                                     | v1.2.0                                            | v1.7.0                             | false  | true             |
| github.com/otiai10/curr                                                     | v1.0.0                                            |                                    | false  | true             |
| github.com/otiai10/mint                                                     | v1.3.1                                            | v1.3.3                             | false  | true             |
| github.com/package-url/packageurl-go                                        | v0.1.1-0.20220203205134-d70459300c8a              |                                    | false  | true             |
| github.com/pact-foundation/pact-go                                          | v1.0.4                                            | v1.7.0                             | false  | true             |
| github.com/pascaldekloe/goe                                                 | v0.1.0                                            |                                    | false  | true             |
| github.com/pborman/uuid                                                     | v1.2.0                                            | v1.2.1                             | false  | true             |
| github.com/pelletier/go-buffruneio                                          | v0.2.0                                            | v0.3.0                             | false  | true             |
| github.com/pelletier/go-toml                                                | v1.9.5                                            |                                    | false  | true             |
| github.com/pelletier/go-toml/v2                                             | v2.0.1                                            | v2.0.2                             | false  | true             |
| github.com/performancecopilot/speed                                         | v3.0.0+incompatible                               |                                    | false  | true             |
| github.com/peterbourgon/diskv                                               | v2.0.1+incompatible                               |                                    | false  | true             |
| github.com/peterh/liner                                                     | v0.0.0-20170211195444-bf27d3ba8e1d                | v1.2.2                             | false  | true             |
| github.com/phayes/checkstyle                                                | v0.0.0-20170904204023-bfd46e6a821d                |                                    | false  | true             |
| github.com/philhofer/fwd                                                    | v1.1.1                                            |                                    | false  | true             |
| github.com/pierrec/lz4                                                      | v2.6.1+incompatible                               |                                    | false  | true             |
| github.com/pierrec/lz4/v4                                                   | v4.1.2                                            | v4.1.15                            | false  | true             |
| github.com/pkg/browser                                                      | v0.0.0-20210911075715-681adbf594b8                |                                    | false  | true             |
| github.com/pkg/diff                                                         | v0.0.0-20210226163009-20ebb0f2a09e                |                                    | false  | true             |
| github.com/pkg/errors                                                       | v0.9.1                                            |                                    | false  | true             |
| github.com/pkg/profile                                                      | v1.2.1                                            | v1.6.0                             | false  | true             |
| github.com/pkg/sftp                                                         | v1.13.1                                           | v1.13.5                            | false  | true             |
| github.com/pmezard/go-difflib                                               | v1.0.0                                            |                                    | false  | true             |
| github.com/polyfloyd/go-errorlint                                           | v0.0.0-20210722154253-910bb7978349                | v1.0.0                             | false  | true             |
| github.com/posener/complete                                                 | v1.2.3                                            |                                    | false  | true             |
| github.com/power-devops/perfstat                                            | v0.0.0-20210106213030-5aafc221ea8c                | v0.0.0-20220216144756-c35f1ee13d7c | false  | true             |
| github.com/pquerna/cachecontrol                                             | v0.1.0                                            |                                    | false  | true             |
| github.com/proglottis/gpgme                                                 | v0.1.1                                            | v0.1.3                             | false  | true             |
| github.com/prometheus/client_golang                                         | v1.12.2                                           |                                    | true   | true             |
| github.com/prometheus/client_model                                          | v0.2.0                                            |                                    | false  | true             |
| github.com/prometheus/common                                                | v0.32.1                                           | v0.36.0                            | false  | true             |
| github.com/prometheus/procfs                                                | v0.7.3                                            |                                    | false  | true             |
| github.com/prometheus/statsd_exporter                                       | v0.21.0                                           | v0.22.7                            | false  | true             |
| github.com/prometheus/tsdb                                                  | v0.7.1                                            | v0.10.0                            | false  | true             |
| github.com/protocolbuffers/txtpbfmt                                         | v0.0.0-20201118171849-f6a6b3f636fc                | v0.0.0-20220608084003-fc78c767cd6a | false  | true             |
| github.com/psampaz/go-mod-outdated                                          | v0.8.0                                            |                                    | true   | true             |
| github.com/pseudomuto/protoc-gen-doc                                        | v1.5.0                                            | v1.5.1                             | false  | true             |
| github.com/pseudomuto/protokit                                              | v0.2.0                                            | v0.2.1                             | false  | true             |
| github.com/quasilyte/go-consistent                                          | v0.0.0-20190521200055-c6f3937de18c                | v0.0.0-20220429160651-4e46040fbc82 | false  | true             |
| github.com/quasilyte/go-ruleguard                                           | v0.3.13                                           | v0.3.16                            | false  | true             |
| github.com/quasilyte/go-ruleguard/dsl                                       | v0.3.10                                           | v0.3.21                            | false  | true             |
| github.com/quasilyte/go-ruleguard/rules                                     | v0.0.0-20210428214800-545e0d2e0bf7                | v0.0.0-20220518160837-796828970844 | false  | true             |
| github.com/quasilyte/regex/syntax                                           | v0.0.0-20200407221936-30656e2c4a95                | v0.0.0-20210819130434-b3f0c404a727 | false  | true             |
| github.com/quobyte/api                                                      | v0.1.8                                            | v1.2.2                             | false  | true             |
| github.com/qur/ar                                                           | v0.0.0-20130629153254-282534b91770                |                                    | false  | true             |
| github.com/rcrowley/go-metrics                                              | v0.0.0-20201227073835-cf1acfcdf475                |                                    | false  | true             |
| github.com/remyoudompheng/bigfft                                            | v0.0.0-20170806203942-52369c62f446                | v0.0.0-20200410134404-eec4a21b6bb0 | false  | true             |
| github.com/rivo/uniseg                                                      | v0.2.0                                            |                                    | false  | true             |
| github.com/robfig/cron/v3                                                   | v3.0.1                                            |                                    | false  | true             |
| github.com/rogpeppe/fastuuid                                                | v1.2.0                                            |                                    | false  | true             |
| github.com/rogpeppe/go-internal                                             | v1.8.1                                            |                                    | false  | true             |
| github.com/rootless-containers/rootlesskit                                  | v1.0.1                                            |                                    | false  | true             |
| github.com/rs/cors                                                          | v1.8.2                                            |                                    | false  | true             |
| github.com/rubiojr/go-vhd                                                   | v0.0.0-20200706105327-02e210299021                | v0.0.0-20200706122120-ccecf6c0760f | false  | true             |
| github.com/russross/blackfriday                                             | v1.6.0                                            |                                    | false  | true             |
| github.com/russross/blackfriday/v2                                          | v2.1.0                                            |                                    | false  | true             |
| github.com/ryancurrah/gomodguard                                            | v1.2.3                                            |                                    | false  | true             |
| github.com/ryanrolds/sqlclosecheck                                          | v0.3.0                                            |                                    | false  | true             |
| github.com/ryanuber/columnize                                               | v2.1.0+incompatible                               | v2.1.2+incompatible                | false  | true             |
| github.com/ryanuber/go-glob                                                 | v1.0.0                                            |                                    | false  | true             |
| github.com/safchain/ethtool                                                 | v0.0.0-20210803160452-9aa261dae9b1                | v0.2.0                             | false  | true             |
| github.com/sagikazarmark/crypt                                              | v0.6.0                                            |                                    | false  | true             |
| github.com/samuel/go-zookeeper                                              | v0.0.0-20190923202752-2cc03de413da                | v0.0.0-20201211165307-7117e9ea2414 | false  | true             |
| github.com/sanposhiho/wastedassign/v2                                       | v2.0.6                                            | v2.0.7                             | false  | true             |
| github.com/saschagrunert/ccli                                               | v1.0.2-0.20200423111659-b68f755cc0f5              |                                    | false  | true             |
| github.com/saschagrunert/go-modiff                                          | v1.3.0                                            | v1.3.1                             | false  | true             |
| github.com/sassoftware/go-rpmutils                                          | v0.1.1                                            | v0.2.0                             | false  | true             |
| github.com/sassoftware/relic                                                | v0.0.0-20210427151427-dfb082b79b74                |                                    | false  | true             |
| github.com/satori/go.uuid                                                   | v1.2.0                                            |                                    | false  | true             |
| github.com/sclevine/agouti                                                  | v3.0.0+incompatible                               |                                    | false  | true             |
| github.com/sclevine/spec                                                    | v1.4.0                                            |                                    | false  | true             |
| github.com/sean-/seed                                                       | v0.0.0-20170313163322-e2103e2c3529                |                                    | false  | true             |
| github.com/sebdah/goldie/v2                                                 | v2.5.3                                            |                                    | false  | true             |
| github.com/seccomp/libseccomp-golang                                        | v0.9.2-0.20220502022130-f33da4d89646              | v0.10.0                            | false  | true             |
| github.com/secure-systems-lab/go-securesystemslib                           | v0.4.0                                            |                                    | false  | true             |
| github.com/securego/gosec                                                   | v0.0.0-20200103095621-79fbf3af8d83                | v0.0.0-20200401082031-e946c8c39989 | false  | true             |
| github.com/securego/gosec/v2                                                | v2.9.1                                            | v2.12.0                            | false  | true             |
| github.com/segmentio/ksuid                                                  | v1.0.4                                            |                                    | false  | true             |
| github.com/sendgrid/rest                                                    | v2.6.9+incompatible                               |                                    | false  | true             |
| github.com/sendgrid/sendgrid-go                                             | v3.11.1+incompatible                              |                                    | false  | true             |
| github.com/sergi/go-diff                                                    | v1.2.0                                            |                                    | false  | true             |
| github.com/shazow/go-diff                                                   | v0.0.0-20160112020656-b6b7b6733b8c                |                                    | false  | true             |
| github.com/shibumi/go-pathspec                                              | v1.3.0                                            |                                    | false  | true             |
| github.com/shirou/gopsutil                                                  | v0.0.0-20190901111213-e4ec7b275ada                | v3.21.11+incompatible              | false  | true             |
| github.com/shirou/gopsutil/v3                                               | v3.22.5                                           | v3.22.6                            | false  | true             |
| github.com/shirou/w32                                                       | v0.0.0-20160930032740-bb4de0191aa4                |                                    | false  | true             |
| github.com/shopspring/decimal                                               | v1.2.0                                            | v1.3.1                             | false  | true             |
| github.com/shurcooL/githubv4                                                | v0.0.0-20220115235240-a14260e6f8a2                | v0.0.0-20220520033151-0b4e3294ff00 | false  | true             |
| github.com/shurcooL/go                                                      | v0.0.0-20180423040247-9e1955d9fb6e                | v0.0.0-20200502201357-93f07166e636 | false  | true             |
| github.com/shurcooL/go-goon                                                 | v0.0.0-20170922171312-37c2f522c041                | v1.0.0                             | false  | true             |
| github.com/shurcooL/graphql                                                 | v0.0.0-20200928012149-18c5c3165e3a                | v0.0.0-20220606043923-3cf50f8a0a29 | false  | true             |
| github.com/shurcooL/sanitized_anchor_name                                   | v1.0.0                                            |                                    | false  | true             |
| github.com/sigstore/cosign                                                  | v1.9.0                                            |                                    | false  | true             |
| github.com/sigstore/fulcio                                                  | v0.1.2-0.20220114150912-86a2036f9bc7              | v0.5.1                             | false  | true             |
| github.com/sigstore/rekor                                                   | v0.4.1-0.20220114213500-23f583409af3              | v0.9.1                             | false  | true             |
| github.com/sigstore/sigstore                                                | v1.2.1-0.20220424143412-3d41663116d5              | v1.3.0                             | false  | true             |
| github.com/sirupsen/logrus                                                  | v1.8.1                                            |                                    | true   | true             |
| github.com/sivchari/tenv                                                    | v1.4.7                                            | v1.6.0                             | false  | true             |
| github.com/skratchdot/open-golang                                           | v0.0.0-20200116055534-eef842397966                |                                    | false  | true             |
| github.com/smallstep/assert                                                 | v0.0.0-20200723003110-82e2b9b3b262                |                                    | false  | true             |
| github.com/smartystreets/assertions                                         | v1.1.0                                            | v1.13.0                            | false  | true             |
| github.com/smartystreets/go-aws-auth                                        | v0.0.0-20180515143844-0c1422d1fdb9                |                                    | false  | true             |
| github.com/smartystreets/goconvey                                           | v1.6.4                                            | v1.7.2                             | false  | true             |
| github.com/smartystreets/gunit                                              | v1.0.0                                            | v1.4.3                             | false  | true             |
| github.com/soheilhy/cmux                                                    | v0.1.5                                            |                                    | true   | true             |
| github.com/sonatard/noctx                                                   | v0.0.1                                            |                                    | false  | true             |
| github.com/songgao/water                                                    | v0.0.0-20200317203138-2b4b6d7c09d8                |                                    | false  | true             |
| github.com/sony/gobreaker                                                   | v0.4.1                                            | v0.5.0                             | false  | true             |
| github.com/sourcegraph/go-diff                                              | v0.6.1                                            |                                    | false  | true             |
| github.com/spaolacci/murmur3                                                | v1.1.0                                            |                                    | false  | true             |
| github.com/spf13/afero                                                      | v1.8.2                                            |                                    | false  | true             |
| github.com/spf13/cast                                                       | v1.5.0                                            |                                    | false  | true             |
| github.com/spf13/cobra                                                      | v1.5.0                                            |                                    | false  | true             |
| github.com/spf13/jwalterweatherman                                          | v1.1.0                                            |                                    | false  | true             |
| github.com/spf13/pflag                                                      | v1.0.5                                            |                                    | false  | true             |
| github.com/spf13/viper                                                      | v1.12.0                                           |                                    | false  | true             |
| github.com/spiegel-im-spiegel/errs                                          | v1.0.5                                            | v1.1.0                             | false  | true             |
| github.com/spiegel-im-spiegel/go-cvss                                       | v1.0.0                                            | v1.1.0                             | false  | true             |
| github.com/spiffe/go-spiffe/v2                                              | v2.1.0                                            | v2.1.1                             | false  | true             |
| github.com/src-d/gcfg                                                       | v1.4.0                                            |                                    | false  | true             |
| github.com/ssgreg/nlreturn/v2                                               | v2.2.1                                            |                                    | false  | true             |
| github.com/stefanberger/go-pkcs11uri                                        | v0.0.0-20201008174630-78d3cae3a980                |                                    | false  | true             |
| github.com/stoewer/go-strcase                                               | v1.2.0                                            |                                    | false  | true             |
| github.com/storageos/go-api                                                 | v2.2.0+incompatible                               | v2.6.0+incompatible                | false  | true             |
| github.com/streadway/amqp                                                   | v1.0.0                                            |                                    | false  | true             |
| github.com/streadway/handy                                                  | v0.0.0-20190108123426-d5acb3125c2a                | v0.0.0-20200128134331-0f66f006fb2e | false  | true             |
| github.com/stretchr/objx                                                    | v0.4.0                                            |                                    | false  | true             |
| github.com/stretchr/testify                                                 | v1.8.0                                            |                                    | true   | true             |
| github.com/subosito/gotenv                                                  | v1.3.0                                            | v1.4.0                             | false  | true             |
| github.com/sylabs/sif/v2                                                    | v2.7.0                                            | v2.7.1                             | false  | true             |
| github.com/sylvia7788/contextcheck                                          | v1.0.4                                            | v1.0.5                             | false  | true             |
| github.com/syndtr/gocapability                                              | v0.0.0-20200815063812-42c35b437635                |                                    | true   | true             |
| github.com/syndtr/goleveldb                                                 | v1.0.1-0.20210819022825-2ae1ddf74ef7              |                                    | false  | true             |
| github.com/tchap/go-patricia                                                | v2.3.0+incompatible                               |                                    | false  | true             |
| github.com/tdakkota/asciicheck                                              | v0.0.0-20200416200610-e657995f937b                | v0.1.1                             | false  | true             |
| github.com/tenntenn/modver                                                  | v1.0.1                                            |                                    | false  | true             |
| github.com/tenntenn/text/transform                                          | v0.0.0-20200319021203-7eef512accb3                |                                    | false  | true             |
| github.com/tent/canonical-json-go                                           | v0.0.0-20130607151641-96e4ba3a7613                |                                    | false  | true             |
| github.com/tetafro/godot                                                    | v1.4.11                                           |                                    | false  | true             |
| github.com/thales-e-security/pool                                           | v0.0.2                                            |                                    | false  | true             |
| github.com/theupdateframework/go-tuf                                        | v0.3.0                                            | v0.3.1                             | false  | true             |
| github.com/tidwall/pretty                                                   | v1.2.0                                            |                                    | false  | true             |
| github.com/tilinna/clock                                                    | v1.1.0                                            |                                    | false  | true             |
| github.com/timakin/bodyclose                                                | v0.0.0-20200424151742-cb6215831a94                | v0.0.0-20210704033933-f49887972144 | false  | true             |
| github.com/tinylib/msgp                                                     | v1.1.5                                            | v1.1.6                             | false  | true             |
| github.com/titanous/rocacheck                                               | v0.0.0-20171023193734-afe73141d399                |                                    | false  | true             |
| github.com/tj/assert                                                        | v0.0.0-20171129193455-018094318fb0                | v0.0.3                             | false  | true             |
| github.com/tj/go-elastic                                                    | v0.0.0-20171221160941-36157cbbebc2                |                                    | false  | true             |
| github.com/tj/go-kinesis                                                    | v0.0.0-20171128231115-08b17f58cb1b                |                                    | false  | true             |
| github.com/tj/go-spin                                                       | v1.1.0                                            |                                    | false  | true             |
| github.com/tklauser/go-sysconf                                              | v0.3.10                                           |                                    | false  | true             |
| github.com/tklauser/numcpus                                                 | v0.4.0                                            | v0.5.0                             | false  | true             |
| github.com/tmc/grpc-websocket-proxy                                         | v0.0.0-20201229170055-e5319fda7802                | v0.0.0-20220101234140-673ab2c3ae75 | false  | true             |
| github.com/tomarrell/wrapcheck/v2                                           | v2.4.0                                            | v2.6.2                             | false  | true             |
| github.com/tomasen/realip                                                   | v0.0.0-20180522021738-f0c99a92ddce                |                                    | false  | true             |
| github.com/tommy-muehle/go-mnd                                              | v1.3.1-0.20200224220436-e6f9a994e8fa              |                                    | false  | true             |
| github.com/tommy-muehle/go-mnd/v2                                           | v2.4.0                                            | v2.5.0                             | false  | true             |
| github.com/transparency-dev/merkle                                          | v0.0.1                                            |                                    | false  | true             |
| github.com/tsenart/vegeta/v12                                               | v12.8.4                                           |                                    | false  | true             |
| github.com/tv42/httpunix                                                    | v0.0.0-20191220191345-2ba4b9c3382c                |                                    | false  | true             |
| github.com/u-root/uio                                                       | v0.0.0-20210528114334-82958018845c                | v0.0.0-20220204230159-dac05f7d2cb4 | false  | true             |
| github.com/uber/jaeger-client-go                                            | v2.30.0+incompatible                              |                                    | false  | true             |
| github.com/ugorji/go                                                        | v1.1.7                                            | v1.2.7                             | false  | true             |
| github.com/ugorji/go/codec                                                  | v1.1.7                                            | v1.2.7                             | false  | true             |
| github.com/ulikunitz/xz                                                     | v0.5.10                                           |                                    | false  | true             |
| github.com/ultraware/funlen                                                 | v0.0.3                                            |                                    | false  | true             |
| github.com/ultraware/whitespace                                             | v0.0.4                                            | v0.0.5                             | false  | true             |
| github.com/urfave/cli                                                       | v1.22.5                                           | v1.22.9                            | false  | true             |
| github.com/urfave/cli/v2                                                    | v2.10.3                                           | v2.11.0                            | true   | true             |
| github.com/urfave/negroni                                                   | v1.0.0                                            |                                    | false  | true             |
| github.com/uudashr/gocognit                                                 | v1.0.5                                            | v1.0.6                             | false  | true             |
| github.com/valyala/bytebufferpool                                           | v1.0.0                                            |                                    | false  | true             |
| github.com/valyala/fasthttp                                                 | v1.30.0                                           | v1.38.0                            | false  | true             |
| github.com/valyala/quicktemplate                                            | v1.7.0                                            |                                    | false  | true             |
| github.com/valyala/tcplisten                                                | v1.0.0                                            |                                    | false  | true             |
| github.com/vbatts/tar-split                                                 | v0.11.2                                           |                                    | false  | true             |
| github.com/vbauerster/mpb/v7                                                | v7.4.1                                            | v7.4.2                             | false  | true             |
| github.com/vektah/gqlparser                                                 | v1.1.2                                            | v1.3.1                             | false  | true             |
| github.com/viki-org/dnscache                                                | v0.0.0-20130720023526-c70c1f23c5d8                |                                    | false  | true             |
| github.com/vishvananda/netlink                                              | v1.2.1-beta.2                                     |                                    | true   | true             |
| github.com/vishvananda/netns                                                | v0.0.0-20210104183010-2eb08e3e575f                | v0.0.0-20211101163701-50045581ed74 | false  | true             |
| github.com/vmihailenco/msgpack/v4                                           | v4.3.12                                           |                                    | false  | true             |
| github.com/vmihailenco/tagparser                                            | v0.1.1                                            | v0.1.2                             | false  | true             |
| github.com/vmware/govmomi                                                   | v0.20.3                                           | v0.29.0                            | false  | true             |
| github.com/weppos/publicsuffix-go                                           | v0.15.1-0.20220329081811-9a40b608a236             |                                    | false  | true             |
| github.com/willf/bitset                                                     | v1.1.11                                           | v1.2.2                             | false  | true             |
| github.com/withfig/autocomplete-tools/packages/cobra                        | v0.0.0-20220122124547-31d3821a6898                | v1.2.0                             | false  | true             |
| github.com/xanzy/go-gitlab                                                  | v0.68.0                                           | v0.68.2                            | false  | true             |
| github.com/xanzy/ssh-agent                                                  | v0.3.0                                            | v0.3.1                             | false  | true             |
| github.com/xdg-go/pbkdf2                                                    | v1.0.0                                            |                                    | false  | true             |
| github.com/xdg-go/scram                                                     | v1.0.2                                            | v1.1.1                             | false  | true             |
| github.com/xdg-go/stringprep                                                | v1.0.2                                            | v1.0.3                             | false  | true             |
| github.com/xdg/scram                                                        | v0.0.0-20180814205039-7eeb5667e42c                | v1.0.5                             | false  | true             |
| github.com/xdg/stringprep                                                   | v0.0.0-20180714160509-73f8eece6fdc                | v1.0.3                             | false  | true             |
| github.com/xeipuuv/gojsonpointer                                            | v0.0.0-20190905194746-02993c407bfb                |                                    | false  | true             |
| github.com/xeipuuv/gojsonreference                                          | v0.0.0-20180127040603-bd5ef7bd5415                |                                    | false  | true             |
| github.com/xeipuuv/gojsonschema                                             | v1.2.0                                            |                                    | false  | true             |
| github.com/xi2/xz                                                           | v0.0.0-20171230120015-48954b6210f8                |                                    | false  | true             |
| github.com/xiang90/probing                                                  | v0.0.0-20190116061207-43a291ad63a2                |                                    | false  | true             |
| github.com/xlab/treeprint                                                   | v0.0.0-20181112141820-a009c3971eca                | v1.1.0                             | false  | true             |
| github.com/xo/terminfo                                                      | v0.0.0-20210125001918-ca9a967f8778                |                                    | false  | true             |
| github.com/xordataexchange/crypt                                            | v0.0.3-0.20170626215501-b2862e3d0a77              |                                    | false  | true             |
| github.com/xrash/smetrics                                                   | v0.0.0-20201216005158-039620a65673                |                                    | false  | true             |
| github.com/yashtewari/glob-intersection                                     | v0.0.0-20180916065949-5c77d914dd0b                | v0.1.0                             | false  | true             |
| github.com/yeya24/promlinter                                                | v0.1.0                                            | v0.2.0                             | false  | true             |
| github.com/youmark/pkcs8                                                    | v0.0.0-20181117223130-1be2e3e5546d                | v0.0.0-20201027041543-1326539a0a0a | false  | true             |
| github.com/ysmood/goob                                                      | v0.4.0                                            |                                    | false  | true             |
| github.com/ysmood/got                                                       | v0.15.1                                           | v0.31.2                            | false  | true             |
| github.com/ysmood/gotrace                                                   | v0.2.2                                            | v0.6.0                             | false  | true             |
| github.com/ysmood/gson                                                      | v0.7.1                                            | v0.7.2                             | false  | true             |
| github.com/ysmood/leakless                                                  | v0.7.0                                            | v0.8.0                             | false  | true             |
| github.com/yudai/gojsondiff                                                 | v1.0.0                                            |                                    | false  | true             |
| github.com/yudai/golcs                                                      | v0.0.0-20170316035057-ecda9a501e82                |                                    | false  | true             |
| github.com/yudai/pp                                                         | v2.0.1+incompatible                               |                                    | false  | true             |
| github.com/yuin/goldmark                                                    | v1.4.12                                           | v1.4.13                            | false  | true             |
| github.com/yusufpapurcu/wmi                                                 | v1.2.2                                            |                                    | false  | true             |
| github.com/yvasiyarov/go-metrics                                            | v0.0.0-20140926110328-57bccd1ccd43                | v0.0.0-20150112132944-c25f46c4b940 | false  | true             |
| github.com/yvasiyarov/gorelic                                               | v0.0.0-20141212073537-a9bba5b9ab50                | v0.0.7                             | false  | true             |
| github.com/yvasiyarov/newrelic_platform_go                                  | v0.0.0-20140908184405-b21fdbd4370f                | v0.0.0-20160601141957-9c099fbc30e9 | false  | true             |
| github.com/zalando/go-keyring                                               | v0.1.1                                            | v0.2.1                             | false  | true             |
| github.com/zeebo/errs                                                       | v1.2.2                                            | v1.3.0                             | false  | true             |
| github.com/zmap/zcrypto                                                     | v0.0.0-20210811211718-6f9bc4aff20f                | v0.0.0-20220701212503-25a10ac913f0 | false  | true             |
| github.com/zmap/zlint/v3                                                    | v3.3.1-0.20211019173530-cb17369b4628              | v3.3.1                             | false  | true             |
| go.etcd.io/bbolt                                                            | v1.3.6                                            |                                    | false  | true             |
| go.etcd.io/etcd                                                             | v0.0.0-20200513171258-e048e166ab9c                | v3.3.27+incompatible               | false  | true             |
| go.etcd.io/etcd/api/v3                                                      | v3.6.0-alpha.0                                    |                                    | false  | true             |
| go.etcd.io/etcd/client/pkg/v3                                               | v3.6.0-alpha.0                                    |                                    | false  | true             |
| go.etcd.io/etcd/client/v2                                                   | v2.306.0-alpha.0                                  |                                    | false  | true             |
| go.etcd.io/etcd/client/v3                                                   | v3.6.0-alpha.0                                    |                                    | false  | true             |
| go.etcd.io/etcd/etcdctl/v3                                                  | v3.6.0-alpha.0                                    |                                    | false  | true             |
| go.etcd.io/etcd/etcdutl/v3                                                  | v3.6.0-alpha.0                                    |                                    | false  | true             |
| go.etcd.io/etcd/pkg/v3                                                      | v3.6.0-alpha.0                                    |                                    | false  | true             |
| go.etcd.io/etcd/raft/v3                                                     | v3.6.0-alpha.0                                    |                                    | false  | true             |
| go.etcd.io/etcd/server/v3                                                   | v3.6.0-alpha.0                                    |                                    | false  | true             |
| go.etcd.io/etcd/tests/v3                                                    | v3.6.0-alpha.0                                    |                                    | false  | true             |
| go.etcd.io/etcd/v3                                                          | v3.6.0-alpha.0.0.20220329071255-af9bd287fe0e      |                                    | false  | true             |
| go.mongodb.org/mongo-driver                                                 | v1.8.3                                            | v1.9.1                             | false  | true             |
| go.mozilla.org/mozlog                                                       | v0.0.0-20170222151521-4bb13139d403                |                                    | false  | true             |
| go.mozilla.org/pkcs7                                                        | v0.0.0-20200128120323-432b2356ecb1                | v0.0.0-20210826202110-33d05740a352 | false  | true             |
| go.opencensus.io                                                            | v0.23.0                                           |                                    | false  | true             |
| go.opentelemetry.io/contrib                                                 | v1.3.0                                            | v1.8.0                             | false  | true             |
| go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc | v0.32.0                                           | v0.33.0                            | true   | true             |
| go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp               | v0.20.0                                           | v0.33.0                            | false  | true             |
| go.opentelemetry.io/contrib/propagators                                     | v0.19.0                                           | v0.22.0                            | false  | true             |
| go.opentelemetry.io/otel                                                    | v1.7.0                                            | v1.8.0                             | true   | true             |
| go.opentelemetry.io/otel/exporters/otlp                                     | v0.20.0                                           |                                    | false  | true             |
| go.opentelemetry.io/otel/exporters/otlp/internal/retry                      | v1.7.0                                            | v1.8.0                             | false  | true             |
| go.opentelemetry.io/otel/exporters/otlp/otlptrace                           | v1.7.0                                            | v1.8.0                             | false  | true             |
| go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc             | v1.7.0                                            | v1.8.0                             | true   | true             |
| go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp             | v1.3.0                                            | v1.8.0                             | false  | true             |
| go.opentelemetry.io/otel/metric                                             | v0.20.0                                           | v0.31.0                            | false  | true             |
| go.opentelemetry.io/otel/oteltest                                           | v0.20.0                                           |                                    | false  | true             |
| go.opentelemetry.io/otel/sdk                                                | v1.7.0                                            | v1.8.0                             | true   | true             |
| go.opentelemetry.io/otel/sdk/export/metric                                  | v0.20.0                                           | v0.28.0                            | false  | true             |
| go.opentelemetry.io/otel/sdk/metric                                         | v0.20.0                                           | v0.31.0                            | false  | true             |
| go.opentelemetry.io/otel/trace                                              | v1.7.0                                            | v1.8.0                             | false  | true             |
| go.opentelemetry.io/proto/otlp                                              | v0.16.0                                           | v0.18.0                            | false  | true             |
| go.starlark.net                                                             | v0.0.0-20200306205701-8dd3e2ee1dd5                | v0.0.0-20220328144851-d1966c6b9fcd | false  | true             |
| go.step.sm/crypto                                                           | v0.14.0                                           | v0.16.2                            | false  | true             |
| go.uber.org/atomic                                                          | v1.9.0                                            |                                    | false  | true             |
| go.uber.org/automaxprocs                                                    | v1.4.0                                            | v1.5.1                             | false  | true             |
| go.uber.org/goleak                                                          | v1.1.12                                           |                                    | false  | true             |
| go.uber.org/multierr                                                        | v1.7.0                                            | v1.8.0                             | false  | true             |
| go.uber.org/tools                                                           | v0.0.0-20190618225709-2cfd321de3ee                |                                    | false  | true             |
| go.uber.org/zap                                                             | v1.21.0                                           |                                    | false  | true             |
| gocloud.dev                                                                 | v0.24.1-0.20211119014450-028788aaaa4c             | v0.25.0                            | false  | true             |
| golang.org/dl                                                               | v0.0.0-20190829154251-82a15e2f2ead                | v0.0.0-20220712162555-2e3db24abc1e | false  | true             |
| golang.org/x/crypto                                                         | v0.0.0-20220411220226-7b82a4e95df4                | v0.0.0-20220622213112-05595931fe9d | false  | true             |
| golang.org/x/exp                                                            | v0.0.0-20210220032938-85be41e4509f                | v0.0.0-20220706164943-b4a6d9510983 | false  | true             |
| golang.org/x/image                                                          | v0.0.0-20190802002840-cff245a6509b                | v0.0.0-20220617043117-41969df76e82 | false  | true             |
| golang.org/x/lint                                                           | v0.0.0-20210508222113-6edffad5e616                |                                    | false  | true             |
| golang.org/x/mobile                                                         | v0.0.0-20201217150744-e6ae53a27f4f                | v0.0.0-20220518205345-8578da9835fd | false  | true             |
| golang.org/x/mod                                                            | v0.6.0-dev.0.20220419223038-86c51ed26bb4          |                                    | false  | true             |
| golang.org/x/net                                                            | v0.0.0-20220621193019-9d032be2e588                | v0.0.0-20220708220712-1185a9018129 | true   | true             |
| golang.org/x/oauth2                                                         | v0.0.0-20220524215830-622c5d57e401                | v0.0.0-20220630143837-2104d58473e0 | false  | true             |
| golang.org/x/sync                                                           | v0.0.0-20220601150217-0de741cfad7f                |                                    | true   | true             |
| golang.org/x/sys                                                            | v0.0.0-20220615213510-4f61da869c0c                | v0.0.0-20220712014510-0a85c31ab51e | true   | true             |
| golang.org/x/term                                                           | v0.0.0-20220526004731-065cf7ba2467                |                                    | false  | true             |
| golang.org/x/text                                                           | v0.3.7                                            |                                    | false  | true             |
| golang.org/x/time                                                           | v0.0.0-20220411224347-583f2d630306                | v0.0.0-20220609170525-579cf78fd858 | false  | true             |
| golang.org/x/tools                                                          | v0.1.11                                           |                                    | false  | true             |
| golang.org/x/xerrors                                                        | v0.0.0-20220517211312-f3a8303e98df                | v0.0.0-20220609144429-65e65417b02f | false  | true             |
| gomodules.xyz/jsonpatch/v2                                                  | v2.2.0                                            |                                    | false  | true             |
| gonum.org/v1/gonum                                                          | v0.6.2                                            | v0.11.0                            | false  | true             |
| gonum.org/v1/netlib                                                         | v0.0.0-20190331212654-76723241ea4e                | v0.0.0-20220323200511-14de99971b2d | false  | true             |
| gonum.org/v1/plot                                                           | v0.0.0-20190515093506-e2840ee46a6b                | v0.11.0                            | false  | true             |
| google.golang.org/api                                                       | v0.82.0                                           | v0.87.0                            | false  | true             |
| google.golang.org/appengine                                                 | v1.6.7                                            |                                    | false  | true             |
| google.golang.org/cloud                                                     | v0.0.0-20151119220103-975617b05ea8                | v0.103.0                           | false  | true             |
| google.golang.org/genproto                                                  | v0.0.0-20220527130721-00d5c0f3be58                | v0.0.0-20220712132514-bdd2acd4974d | false  | true             |
| google.golang.org/grpc                                                      | v1.47.0                                           | v1.48.0                            | true   | true             |
| google.golang.org/grpc/cmd/protoc-gen-go-grpc                               | v1.1.0                                            | v1.2.0                             | false  | true             |
| google.golang.org/grpc/examples                                             | v0.0.0-20201130180447-c456688b1860                | v0.0.0-20220712211207-5e15eac0c4df | false  | true             |
| google.golang.org/protobuf                                                  | v1.28.0                                           |                                    | false  | true             |
| gopkg.in/airbrake/gobrake.v2                                                | v2.0.9                                            |                                    | false  | true             |
| gopkg.in/alecthomas/kingpin.v2                                              | v2.2.6                                            |                                    | false  | true             |
| gopkg.in/alexcesaro/statsd.v2                                               | v2.0.0                                            |                                    | false  | true             |
| gopkg.in/check.v1                                                           | v1.0.0-20201130134442-10cb98267c6c                |                                    | false  | true             |
| gopkg.in/cheggaaa/pb.v1                                                     | v1.0.28                                           |                                    | false  | true             |
| gopkg.in/errgo.v2                                                           | v2.1.0                                            |                                    | false  | true             |
| gopkg.in/fsnotify.v1                                                        | v1.4.7                                            |                                    | false  | true             |
| gopkg.in/gcfg.v1                                                            | v1.2.3                                            |                                    | false  | true             |
| gopkg.in/gemnasium/logrus-airbrake-hook.v2                                  | v2.1.2                                            |                                    | false  | true             |
| gopkg.in/go-playground/assert.v1                                            | v1.2.1                                            |                                    | false  | true             |
| gopkg.in/go-playground/validator.v9                                         | v9.29.1                                           | v9.31.0                            | false  | true             |
| gopkg.in/inf.v0                                                             | v0.9.1                                            |                                    | false  | true             |
| gopkg.in/ini.v1                                                             | v1.66.4                                           | v1.66.6                            | false  | true             |
| gopkg.in/natefinch/lumberjack.v2                                            | v2.0.0                                            |                                    | false  | true             |
| gopkg.in/resty.v1                                                           | v1.12.0                                           |                                    | false  | true             |
| gopkg.in/square/go-jose.v2                                                  | v2.6.0                                            |                                    | false  | true             |
| gopkg.in/src-d/go-billy.v4                                                  | v4.3.2                                            |                                    | false  | true             |
| gopkg.in/src-d/go-git-fixtures.v3                                           | v3.5.0                                            |                                    | false  | true             |
| gopkg.in/src-d/go-git.v4                                                    | v4.13.1                                           |                                    | false  | true             |
| gopkg.in/tomb.v1                                                            | v1.0.0-20141024135613-dd632973f1e7                |                                    | false  | true             |
| gopkg.in/warnings.v0                                                        | v0.1.2                                            |                                    | false  | true             |
| gopkg.in/yaml.v2                                                            | v2.4.0                                            |                                    | false  | true             |
| gopkg.in/yaml.v3                                                            | v3.0.1                                            |                                    | false  | true             |
| gotest.tools                                                                | v2.2.0+incompatible                               |                                    | false  | true             |
| gotest.tools/v3                                                             | v3.0.3                                            | v3.3.0                             | false  | true             |
| honnef.co/go/tools                                                          | v0.2.1                                            | v0.3.2                             | false  | true             |
| k8s.io/api                                                                  | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | true   | true             |
| k8s.io/apiextensions-apiserver                                              | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/apimachinery                                                         | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | true   | true             |
| k8s.io/apiserver                                                            | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/cli-runtime                                                          | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/client-go                                                            | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | true   | true             |
| k8s.io/cloud-provider                                                       | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/cluster-bootstrap                                                    | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/code-generator                                                       | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/component-base                                                       | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/component-helpers                                                    | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/controller-manager                                                   | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/cri-api                                                              | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | true   | true             |
| k8s.io/csi-translation-lib                                                  | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/gengo                                                                | v0.0.0-20220307231824-4627b89bbf1b                | v0.0.0-20220613173612-397b4ae3bce7 | false  | true             |
| k8s.io/klog                                                                 | v1.0.0                                            |                                    | false  | true             |
| k8s.io/klog/v2                                                              | v2.70.1                                           |                                    | true   | true             |
| k8s.io/kube-aggregator                                                      | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/kube-controller-manager                                              | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/kube-openapi                                                         | v0.0.0-20220603121420-31174f50af60                | v0.0.0-20220627174259-011e075b9cb8 | false  | true             |
| k8s.io/kube-proxy                                                           | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/kube-scheduler                                                       | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/kubectl                                                              | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/kubelet                                                              | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/kubernetes                                                           | v1.25.0-alpha.1.0.20220624161656-6219eed24f3c     | v1.25.0-alpha.2                    | true   | true             |
| k8s.io/legacy-cloud-providers                                               | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/metrics                                                              | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/mount-utils                                                          | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/pod-security-admission                                               | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/release                                                              | v0.14.0                                           |                                    | true   | true             |
| k8s.io/sample-apiserver                                                     | v0.0.0-20220624161656-6219eed24f3c                | v0.0.0-20220712183556-71481bf24746 | false  | true             |
| k8s.io/system-validators                                                    | v1.7.0                                            |                                    | false  | true             |
| k8s.io/utils                                                                | v0.0.0-20220210201930-3a6ce19ff2f9                | v0.0.0-20220706174534-f6158b442e7c | true   | true             |
| knative.dev/hack                                                            | v0.0.0-20220224013837-e1785985d364                | v0.0.0-20220701014203-65c463ac8c98 | false  | true             |
| knative.dev/hack/schema                                                     | v0.0.0-20220224013837-e1785985d364                | v0.0.0-20220701014203-65c463ac8c98 | false  | true             |
| knative.dev/pkg                                                             | v0.0.0-20220325200448-1f7514acd0c2                | v0.0.0-20220705130606-e60d250dc637 | false  | true             |
| modernc.org/cc                                                              | v1.0.0                                            | v1.0.1                             | false  | true             |
| modernc.org/golex                                                           | v1.0.0                                            | v1.0.1                             | false  | true             |
| modernc.org/mathutil                                                        | v1.0.0                                            | v1.4.1                             | false  | true             |
| modernc.org/strutil                                                         | v1.0.0                                            | v1.1.2                             | false  | true             |
| modernc.org/xc                                                              | v1.0.0                                            |                                    | false  | true             |
| mvdan.cc/editorconfig                                                       | v0.2.0                                            |                                    | false  | true             |
| mvdan.cc/gofumpt                                                            | v0.1.1                                            | v0.3.1                             | false  | true             |
| mvdan.cc/interfacer                                                         | v0.0.0-20180901003855-c20040233aed                |                                    | false  | true             |
| mvdan.cc/lint                                                               | v0.0.0-20170908181259-adc824a0674b                |                                    | false  | true             |
| mvdan.cc/sh/v3                                                              | v3.5.1                                            |                                    | true   | true             |
| mvdan.cc/unparam                                                            | v0.0.0-20210104141923-aac4ce9116a7                | v0.0.0-20220706161116-678bad134442 | false  | true             |
| nhooyr.io/websocket                                                         | v1.8.7                                            |                                    | false  | true             |
| pack.ag/amqp                                                                | v0.11.2                                           | v0.12.5                            | false  | true             |
| rsc.io/binaryregexp                                                         | v0.2.0                                            |                                    | false  | true             |
| rsc.io/pdf                                                                  | v0.1.1                                            |                                    | false  | true             |
| rsc.io/quote/v3                                                             | v3.1.0                                            |                                    | false  | true             |
| rsc.io/sampler                                                              | v1.3.0                                            | v1.99.99                           | false  | true             |
| sigs.k8s.io/apiserver-network-proxy/konnectivity-client                     | v0.0.32                                           |                                    | false  | true             |
| sigs.k8s.io/bom                                                             | v0.3.0-rc1.0.20220627190259-3850e8ff14fa          | v0.3.0                             | true   | true             |
| sigs.k8s.io/json                                                            | v0.0.0-20211208200746-9f7c6b3444d2                | v0.0.0-20220525155127-227cbc7cc124 | false  | true             |
| sigs.k8s.io/kustomize/api                                                   | v0.11.4                                           | v0.11.5                            | false  | true             |
| sigs.k8s.io/kustomize/cmd/config                                            | v0.10.6                                           | v0.10.7                            | false  | true             |
| sigs.k8s.io/kustomize/kustomize/v4                                          | v4.5.4                                            | v4.5.5                             | false  | true             |
| sigs.k8s.io/kustomize/kyaml                                                 | v0.13.6                                           | v0.13.7                            | false  | true             |
| sigs.k8s.io/mdtoc                                                           | v1.1.0                                            |                                    | false  | true             |
| sigs.k8s.io/promo-tools/v3                                                  | v3.4.2                                            | v3.4.4                             | false  | true             |
| sigs.k8s.io/release-sdk                                                     | v0.9.1                                            |                                    | true   | true             |
| sigs.k8s.io/release-utils                                                   | v0.7.1                                            |                                    | true   | true             |
| sigs.k8s.io/structured-merge-diff/v4                                        | v4.2.1                                            |                                    | false  | true             |
| sigs.k8s.io/yaml                                                            | v1.3.0                                            |                                    | true   | true             |
| sigs.k8s.io/zeitgeist                                                       | v0.3.0                                            |                                    | true   | true             |
| sourcegraph.com/sourcegraph/appdash                                         | v0.0.0-20190731080439-ebfcffb1b5c0                | v0.0.0-20211028080628-e2786a622600 | false  | true             |
| sourcegraph.com/sqs/pbtypes                                                 | v0.0.0-20180604144634-d3ebe8f20ae4                | v1.0.0                             | false  | true             |
