---
title: "GDPR"
description: "General Data Protection Regulation"
weight: 130
menu:
  docs:
    parent: "compliancy"
---

# The EU General Data Protection Regulation (GDPR)

The EU General Data Protection Regulation (GDPR) replaces the Data Protection Directive 95/46/EC and was designed to harmonize data privacy laws across Europe, to protect and empower all EU citizens&#39; 
data privacy and to reshape the way organizations across the region approach data privacy.

* **Requirement**  NLX0051 
  * **Source**  EU General Data Protection Regulation (GDPR) 
  * **Category**  Right to Access 
  * **Type**  Mandatory 
  * **Compliant**  Yes 
  * **Description**  Part of the expanded rights of data subjects outlined by the GDPR is the right for data subjects to obtain from the data controller confirmation as to whether or not personal data concerning them is being processed, where and for what purpose. Further, the controller shall provide a copy of the personal data, free of charge, in an electronic format. This change is a dramatic shift to data transparency and empowerment of data subjects. 
  * **Implications** 
     * Although personal data is processed through the APIs published on NLX none of that data is stored by NLX;
      The right to access data used by the API&#39;s will be offered to subjects by the API provider and not NLX;
     * NLX maintains transaction logs. These logs contain personal data if API&#39;s have been accessed that access personal data. If this is the case then the _Right to Access_ applies to this data applies. NLX will provide means for citizens to access their personal data in an electronic format.
* **Requirement**  NLX0052 
  * **Source**  EU General Data Protection Regulation (GDPR) 
  * **Category**  Right to be Forgotten 
  * **Type**  Mandatory 
  * **Compliant**  N/A
  * **Description**  Also known as Data Erasure, the right to be forgotten entitles the data subject to have the data controller erase his/her personal data, cease further dissemination of the data, and potentially have third parties halt processing of the data. The conditions for erasure, as outlined in article 17, include the data no longer being relevant to original purposes for processing, or a data subjects withdrawing consent. It should also be noted that this right requires controllers to compare the subjects&#39; rights to &quot;the public interest in the availability of the data&quot; when considering such requests.
  * **Implications** Although personal data is processed through the APIs published on NLX none of that data is stored by NLX;
     * NLX maintains transaction logs. These logs contain personal data if API&#39;s have been accessed that 
      access personal data. The transaction logs are used to provide a way to audit trail the transactions in 
      the federated network. The integrity of these logs is paramount for these audit trails. 
      The _Right to be Forgotten_ will therefore not apply to the NLX transaction logs.
* **Requirement**  NLX0053 
  * **Source**  EU General Data Protection Regulation (GDPR) 
  * **Category**  Data Portability 
  * **Type**  Mandatory 
  * **Compliant**  N/A 
  * **Description**  GDPR introduces data portability - the right for a data subject to receive the personal data concerning them, which they have previously provided in a &#39;commonly use and machine readable format&#39; and have the right to transmit that data to another controller. 
  * **Implications** 
      * Although personal data is processed through the APIs published on NLX none of that data is stored by NLX
      * NLX maintains transaction logs. These logs contain personal data if API&#39;s have been accessed that access personal data.  
      The _Data Portability_ right does not apply to the transaction log records. These records depict service calls which have been handled by the service provider. While in some cases these log records can be related to persons they do not constitute data which is transferrable to other parties.
* **Requirement**  NLX0054 
  *  **Source**  EU General Data Protection Regulation (GDPR) 
  *  **Category**  Privacy by Design 
  *  **Type**  Mandatory 
  *  **Compliant**  Yes 
  *  **Description**  Privacy by design as a concept has existed for years now, but it is only just becoming part of a legal requirement with the GDPR. At its core, privacy by design calls for the inclusion of data protection from the onset of the designing of systems, rather than an addition. More specifically - &#39;The controller shall..implement appropriate technical and organisational measures..in an effective way.. in order to meet the requirements of this Regulation and protect the rights of data subjects&#39;. Article 23 calls for controllers to hold and process only the data absolutely necessary for the completion of its duties (data minimisation), as well as limiting the access to personal data to those needing to act out the processing. 
  *  **Implications** 
     * NLX is designed and build based on the following privacy by design principles
        * Minimize: Limit as much as possible the processing of personal data
        * Separate: Distribute or isolate personal data as much as possible, to prevent Correlation
        * Abstract: Limit as much as possible the detail in which personal data is Processed
        * Hide: Prevent personal data to become public or known
        * Inform: Inform data subjects about the processing of their personal data
        * Control: Provide data subjects control about the processing of their personal data
        * Enforce: Commit to processing personal data in a privacy friendly way, and enforce this
        * Demonstrate: Demonstrate you are processing personal data in a privacy friendly way.
        * NLX does not store any (personal) data, except for identifying numbers in transaction logs. 
          Data stored in the transaction logs is limited to the data which is needed for audit trailing.
