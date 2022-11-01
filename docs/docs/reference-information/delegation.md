---
id: delegation
title: Delegation
---

Delegation is when organization A (delegatee) carries out a task on behalf of organization B (delegator), e.g. based on a Data Processing Agreement (DPA). Most of the times this task involves organization A retrieving information from organization X, Y or Z on behalf of organization B.



Every organization connecting to NLX always uses it's own PKI-Overheid certificate(s) for every action within the NLX network. This way it is always clear which organization is actually connecting, providing a service and / or using a service.

Organizations can delegate access rights to other organizations using an 'order', which enables delegated use of services.  

Note: currently the common practise is to hand over PKI certificates so organizations can impersonate other organizations to make calls 'in the name off' is highly unwanted and unnecessary thanks to this feature.

### The order
NLX provides a digital representation of the DPA, named 'order'. An order contains the scope of the DPA: the duration of the DPA and the services necessary to perform the tasks mentioned in the DPA. The delegator can only choose services to which access already has been granted. The order also contains the reference to the delegator and delegatee and an unique reference to itself.

### The flow
1. The delegator creates the order.
1. The client of the delegatee is configured so it will use the correct order reference and delegator-reference
1. The Outway of the delegatee recognized that the call is on behalf of a delegator and so it first fetches (when not available) a valid claim signed by the delegator
1. The delegator will return a valid claim when the API is part of the order and the order is still valid and the delegatee is the same as on the order
1. The delegatee will send the request to the API including the valid claim
1. The Inway of the API owner will accept this request based on the access of the organization who signed the claim. Depending on this access a response is returned


### Conclusion

When using NLX:
- Delegation adds transparency about which organization is performing which activity
- Organizations do not need information about DPA's of other organizations
- PKI-overheid certificates stay with the owner

### Animated explanation
Click on this [link](https://commonground.nl/files/view/9115e0b9-b80b-494f-9169-252410b0c4cb/multitenant-animatie-met-uitleg.pptx) to open a powerpoint presentation
