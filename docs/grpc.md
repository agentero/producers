# gRPC Documentation




This page was automatically generated based on the protobuf file `producer.proto`.

Paths for the REST proxy of the gRPC interface can be found [here](https://github.com/agentero/producers/blob/main/proto/rest-annotations.yaml).


## producer.ProducerService
Manages producers and their compliance information.

### Methods
#### GetAgencyProducers

GetAgencyProducer returns the list of producers of an agency.

| Request | Response |
| ------- | -------- |
| [`GetAgencyProducersRequest`](#producer.GetAgencyProducersRequest) | [`GetAgencyProducesResponse`](#producer.GetAgencyProducesResponse) |

#### GetProducer

GetProducer returns a producer with the given ID.

| Request | Response |
| ------- | -------- |
| [`GetProducerRequest`](#producer.GetProducerRequest) | [`GetProducerResponse`](#producer.GetProducerResponse) |

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

#### <div id="producer.AgencyLookup">AgencyLookup</div>
Producer lookup by the agency's tax number.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tax_number` | [`string`](#string) |  |  |





#### <div id="producer.GetAgencyProducersRequest">GetAgencyProducersRequest</div>



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `agency_name` | [`string`](#string) |  |  |





#### <div id="producer.GetAgencyProducesResponse">GetAgencyProducesResponse</div>



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `producers` | [`Producer`](#producer.Producer) | repeated |  |





#### <div id="producer.GetProducerRequest">GetProducerRequest</div>
Request to retrieve a producer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `producer_id_lookup` | [`GetProducerRequest.ProducerIDLookup`](#producer.GetProducerRequest.ProducerIDLookup) |  |  |
| `npn_lookup` | [`GetProducerRequest.ProducerNPNLookup`](#producer.GetProducerRequest.ProducerNPNLookup) |  |  |
| `email_lookup` | [`GetProducerRequest.EmailLookup`](#producer.GetProducerRequest.EmailLookup) |  |  |





#### <div id="producer.GetProducerRequest.EmailLookup">GetProducerRequest.EmailLookup</div>



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `email` | [`string`](#string) |  |  |





#### <div id="producer.GetProducerRequest.ProducerIDLookup">GetProducerRequest.ProducerIDLookup</div>



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `producer_id` | [`string`](#string) |  |  |





#### <div id="producer.GetProducerRequest.ProducerNPNLookup">GetProducerRequest.ProducerNPNLookup</div>



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `producer_npn` | [`string`](#string) |  |  |





#### <div id="producer.GetProducerResponse">GetProducerResponse</div>
Response containing a producer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `producer` | [`Producer`](#producer.Producer) |  |  |





#### <div id="producer.License">License</div>
Producers license information and status.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `license_number` | [`string`](#string) |  |  |
| `license_state` | [`string`](#string) |  |  |
| `residency_status` | [`string`](#string) |  |  |
| `active` | [`bool`](#bool) |  |  |
| `status` | [`License.LicenseStatus`](#producer.License.LicenseStatus) |  |  |
| `expiration_date` | [`google.protobuf.Timestamp`](#google.protobuf.Timestamp) |  |  |
| `updated_at` | [`google.protobuf.Timestamp`](#google.protobuf.Timestamp) |  |  |





#### <div id="producer.LicenseLookup">LicenseLookup</div>
Producer lookup by license number and state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `license_number` | [`string`](#string) |  |  |
| `state` | [`string`](#string) |  |  |





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





#### <div id="producer.NPNLookup">NPNLookup</div>
Producer lookup by NPN.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `npn` | [`string`](#string) |  |  |





#### <div id="producer.NewProducer">NewProducer</div>
Representation of a producer that we want to add.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `npn_lookup` | [`NPNLookup`](#producer.NPNLookup) |  |  |
| `agency_lookup` | [`AgencyLookup`](#producer.AgencyLookup) |  |  |
| `license_lookup` | [`LicenseLookup`](#producer.LicenseLookup) |  |  |
| `ssn_lookup` | [`SSNLookup`](#producer.SSNLookup) |  |  |
| `name` | [`string`](#string) |  | The name of the producer. Not required. |
| `email` | [`string`](#string) |  | The email of the producer. Not required. |





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
| `biographic` | [`Producer.Biographic`](#producer.Producer.Biographic) |  | Biographic information of the producer. |
| `regulatory_info` | [`Producer.ProducerRegulatoryInfo`](#producer.Producer.ProducerRegulatoryInfo) |  | Producer's regulatory information. |
| `appointments` | [`Producer.Appointment`](#producer.Producer.Appointment) | repeated | The list of carrier appointments of the producer. |





#### <div id="producer.Producer.Appointment">Producer.Appointment</div>



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `branch_id` | [`string`](#string) |  |  |
| `company_name` | [`string`](#string) |  |  |
| `fein` | [`string`](#string) |  |  |
| `co_code` | [`string`](#string) |  |  |
| `line_of_authority` | [`string`](#string) |  |  |
| `loa_code` | [`string`](#string) |  |  |
| `status` | [`string`](#string) |  |  |
| `termination_reason` | [`string`](#string) |  |  |
| `status_reason_date` | [`string`](#string) |  |  |
| `appointment_renewal_date` | [`google.protobuf.Timestamp`](#google.protobuf.Timestamp) |  |  |
| `agency_affiliations` | [`string`](#string) |  |  |





#### <div id="producer.Producer.Biographic">Producer.Biographic</div>
Biographic information of the producer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `last_name` | [`string`](#string) |  |  |
| `first_name` | [`string`](#string) |  |  |
| `middle_name` | [`string`](#string) |  |  |
| `date_of_birth` | [`google.protobuf.Timestamp`](#google.protobuf.Timestamp) |  |  |
| `ssn` | [`string`](#string) |  |  |
| `fein` | [`string`](#string) |  |  |
| `company_name` | [`string`](#string) |  |  |
| `state_domicile` | [`string`](#string) |  |  |





#### <div id="producer.Producer.ProducerRegulatoryInfo">Producer.ProducerRegulatoryInfo</div>
ProducerRegulatoryInfo contains regulatory information about a producer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `regulatory_actions_by_state` | [`Producer.ProducerRegulatoryInfo.RegulatoryActionsByStateEntry`](#producer.Producer.ProducerRegulatoryInfo.RegulatoryActionsByStateEntry) | repeated | Regulatory actions by state, if any. |
| `clearance_certification_info` | [`string`](#string) |  |  |
| `nasd_exam_details` | [`string`](#string) |  |  |





#### <div id="producer.Producer.ProducerRegulatoryInfo.RegulatoryAction">Producer.ProducerRegulatoryInfo.RegulatoryAction</div>



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `action_id` | [`string`](#string) |  |  |
| `origin_of_action` | [`string`](#string) |  |  |
| `reason_for_action` | [`string`](#string) |  |  |
| `disposition` | [`string`](#string) |  |  |
| `date_of_action` | [`google.protobuf.Timestamp`](#google.protobuf.Timestamp) |  |  |
| `effective_date` | [`google.protobuf.Timestamp`](#google.protobuf.Timestamp) |  |  |
| `enter_date` | [`google.protobuf.Timestamp`](#google.protobuf.Timestamp) |  |  |
| `file_ref` | [`string`](#string) |  |  |
| `penalty_fine_forfeiture` | [`string`](#string) |  |  |
| `length_of_order` | [`string`](#string) |  |  |





#### <div id="producer.Producer.ProducerRegulatoryInfo.RegulatoryActionsByStateEntry">Producer.ProducerRegulatoryInfo.RegulatoryActionsByStateEntry</div>



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [`string`](#string) |  |  |
| `value` | [`Producer.ProducerRegulatoryInfo.RegulatoryAction`](#producer.Producer.ProducerRegulatoryInfo.RegulatoryAction) |  |  |





#### <div id="producer.SSNLookup">SSNLookup</div>
Producer lookup with a matching tax number and last name.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tax_number` | [`string`](#string) |  |  |
| `last_name` | [`string`](#string) |  |  |






### Enums


<a name="producer.License.LicenseStatus"></a>

#### License.LicenseStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| LICENSE_STATUS_UNSPECIFIED | 0 |  |
| LICENSE_STATUS_EXPIRED | 1 |  |
| LICENSE_STATUS_VALID | 2 |  |
| LICENSE_STATUS_NOT_ACTIVE | 3 |  |



