# gRPC Documentation




This page was automatically generated based on the protobuf file `producer.proto`.

Paths for the REST proxy of the gRPC interface can be found [here](https://github.com/agentero/producers/blob/main/proto/rest-annotations.yaml).


## producer.ProducerService
ProducerService defines the API for managing producers.

### Methods
#### ListNewProducers

Return the list of new producers of an Agency. A producer is considered new if: 1. It is not in the list of processed producers. 2. Since its creation, it has not been updated.

| Request | Response |
| ------- | -------- |
| [`ListNewProducersRequest`](#producer.ListNewProducersRequest) | [`ListNewProducersResponse`](#producer.ListNewProducersResponse) |

#### ListUpdatedProducers

Return the list of updated producers of an Agency. A producer is considered updated if: 1. It is not in the list of processed producers. 2. Since its creation, it has been updated.

| Request | Response |
| ------- | -------- |
| [`ListUpdatedProducersRequest`](#producer.ListUpdatedProducersRequest) | [`ListUpdatedProducersResponse`](#producer.ListUpdatedProducersResponse) |

#### MarkAsProcessed

Mark a list of producers as processed, so they will not be returned in future calls to ListNewProducers or ListUpdatedProducers.

| Request | Response |
| ------- | -------- |
| [`MarkAsProcessedRequest`](#producer.MarkAsProcessedRequest) | [`.google.protobuf.Empty`](#google.protobuf.Empty) |

#### NewProducer

NewProducer starts the onboarding process for the given producer.

| Request | Response |
| ------- | -------- |
| [`NewProducerRequest`](#producer.NewProducerRequest) | [`.google.protobuf.Empty`](#google.protobuf.Empty) |

#### NewProducers

NewProducers does the same as NewProducer but accepts a list of producers to onboard.

| Request | Response |
| ------- | -------- |
| [`NewProducersRequest`](#producer.NewProducersRequest) | [`.google.protobuf.Empty`](#google.protobuf.Empty) |




### Messages

#### <div id="producer.License">License</div>



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `license_number` | [`string`](#string) |  |  |
| `license_state` | [`string`](#string) |  |  |
| `residency_status` | [`string`](#string) |  |  |
| `active` | [`bool`](#bool) |  |  |
| `status` | [`License.LicenseStatus`](#producer.License.LicenseStatus) |  |  |
| `expiration_date` | [`google.protobuf.Timestamp`](#google.protobuf.Timestamp) |  |  |
| `updated_at` | [`google.protobuf.Timestamp`](#google.protobuf.Timestamp) |  |  |





#### <div id="producer.ListNewProducersRequest">ListNewProducersRequest</div>
Request to list new producers for an Agency.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `agency_name` | [`string`](#string) | optional |  |





#### <div id="producer.ListNewProducersResponse">ListNewProducersResponse</div>
Response containing a list of new producers for an Agency.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `new_producers` | [`Producer`](#producer.Producer) | repeated |  |





#### <div id="producer.ListUpdatedProducersRequest">ListUpdatedProducersRequest</div>
Request to list updated producers for an Agency.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `agency_name` | [`string`](#string) | optional |  |





#### <div id="producer.ListUpdatedProducersResponse">ListUpdatedProducersResponse</div>
Response containing a list of updated producers for an Agency.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `updated_producers` | [`Producer`](#producer.Producer) | repeated |  |





#### <div id="producer.MarkAsProcessedRequest">MarkAsProcessedRequest</div>
Request to mark a list of producers as processed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `agency_name` | [`string`](#string) |  |  |
| `producer_ids` | [`string`](#string) | repeated |  |





#### <div id="producer.NewProducer">NewProducer</div>
NewProducer is a Producer that has not been processed yet.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `npn` | [`string`](#string) |  |  |
| `name` | [`string`](#string) |  |  |
| `email` | [`string`](#string) |  |  |





#### <div id="producer.NewProducerRequest">NewProducerRequest</div>
Request to onboard a new producer and automatically associate it 
with an Agency.
This will trigger a call to the NIPR API to retrieve license information 
of the producer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `agency_name` | [`string`](#string) |  |  |
| `agency_fein` | [`string`](#string) |  |  |
| `producer` | [`NewProducer`](#producer.NewProducer) |  |  |





#### <div id="producer.NewProducersRequest">NewProducersRequest</div>
Request to onboard a list of producers and automatically associate them 
with an Agency.
This will trigger a call to the NIPR API to retrieve license information of 
every producer in the list.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `agency_name` | [`string`](#string) |  |  |
| `agency_fein` | [`string`](#string) |  |  |
| `producers` | [`NewProducer`](#producer.NewProducer) | repeated |  |





#### <div id="producer.Producer">Producer</div>
Producer represents a producer that has been onboarded.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [`string`](#string) |  | Internal ID of the producer. |
| `name` | [`string`](#string) |  | The name of the producer. |
| `email` | [`string`](#string) |  | The email of the producer. |
| `npn` | [`string`](#string) |  | The NPN of the producer. This is used to retrieve the license information of the producer from the NIPR API. |
| `agency_name` | [`string`](#string) |  |  |
| `licenses` | [`License`](#producer.License) | repeated | The licenses of the producer. |






### Enums


<a name="producer.License.LicenseStatus"></a>

#### License.LicenseStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| LICENSE_STATUS_UNSPECIFIED | 0 |  |
| LICENSE_STATUS_EXPIRED | 1 |  |
| LICENSE_STATUS_VALID | 2 |  |
| LICENSE_STATUS_NOT_ACTIVE | 3 |  |


<a name="producer.LineOfAuthority"></a>

#### LineOfAuthority


| Name | Number | Description |
| ---- | ------ | ----------- |
| LINE_OF_AUTHORITY_UNSPECIFIED | 0 |  |
| LINE_OF_AUTHORITY_LIFE | 1 |  |
| LINE_OF_AUTHORITY_PERSONAL | 2 |  |
| LINE_OF_AUTHORITY_PERSONAL_AUTO | 3 |  |
| LINE_OF_AUTHORITY_PROPERTY_AND_CASUALTY | 4 |  |
| LINE_OF_AUTHORITY_HEALTH | 5 |  |
| LINE_OF_AUTHORITY_COMMERCIAL | 6 |  |
| LINE_OF_AUTHORITY_LIMITED_COMMERCIAL | 7 |  |
| LINE_OF_AUTHORITY_SURPLUS_LINES | 8 |  |
| LINE_OF_AUTHORITY_OTHER | 9 |  |



