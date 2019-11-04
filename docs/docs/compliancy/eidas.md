---
id: eidas
title: eIDAS
---

The eIDAS directive provides for secure and protected data exchange during the delivery of public services. It ensures that people and businesses can use their own national electronic identification schemes (eIDs) to access public services in other EU countries where eIDs are available and creates a European internal market for eTS - namely electronic signatures, electronic seals, time stamp, electronic delivery service and website authentication - by ensuring that they will work across borders and have the same legal status as traditional paper based processes.

The eIDAS regulation:

- lays down the conditions under which Member States recognise electronic identification means of natural and legal persons falling under a notified electronic identification scheme of another Member State;
- lays down rules for trust services, in particular for electronic transactions; and
- establishes a legal framework for electronic signatures, electronic seals, electronic time stamps, electronic documents, electronic registered delivery services and certificate services for website authentication.

This eIDAS Regulation applies to electronic identification schemes that have been notified by a Member State,
and to trust service providers that are established in the European Union. The Regulation does not apply to
the provision of trust services that are used exclusively within closed systems resulting from national law or
from agreements between a defined set of participants and does not affect national or Union law related to
the conclusion and validity of contracts or other legal or procedural obligations relating to form.

* **Requirement**  NLX0048
  * **Source**  eIDAS
  * **Category**  article 8: Assurance levels of electronic identification schemes
  * **Type**  Mandatory
  * **Compliant**  Yes
  * **Description**  An electronic identification scheme notified pursuant to Article 9(1) shall specify assurance levels low, substantial and/or high for electronic identification means issued under that scheme.
                     The assurance levels low, substantial and high shall meet respectively the following criteria:
     * assurance level **low** shall refer to an electronic identification means in the context of an electronic identification scheme, which provides a limited degree of confidence in the claimed or asserted identity of a person, and is characterised with reference to technical specifications, standards and procedures related thereto, including technical controls, the purpose of which is to decrease the risk of misuse or alteration of the identity;
     * assurance level **substantial** shall refer to an electronic identification means in the context of an electronic identification scheme, which provides a substantial degree of confidence in the claimed or asserted identity of a person, and is characterised with reference to technical specifications, standards and procedures related thereto, including technical controls, the purpose of which is to decrease substantially the risk of misuse or alteration of the identity;
     * assurance level **high** shall refer to an electronic identification means in the context of an electronic identification scheme, which provides a higher degree of confidence in the claimed or asserted identity of a person than electronic identification means with the assurance level substantial, and is characterised with reference to technical specifications, standards and procedures related thereto, including technical controls, the purpose of which is to prevent misuse or alteration of the identity.
  * **Implications**
     * NLX will have to support configurations which enables using different assurance levels for APIs.
* **Requirement**  NLX0049
  * **Source**  eIDAS
  * **Category**  article 15: Accessibility for persons with disabilities
  * **Type**  Mandatory
  * **Compliant**  Yes
  * **Description**  Where feasible, trust services provided and end-user products used in the provision of those services shall be made accessible for persons with disabilities.
  * **Implications**
     * The API Discovery user interface which publishes the NLX API&#39;s must comply with &quot;WCAG 2.0 Success Criterion 1.1.1 Non-text content&quot; as specified in the EN 301 549 v1.1.2 standard.
