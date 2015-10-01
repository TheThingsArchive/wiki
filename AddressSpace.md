Until there is an official API for registering blocks
of nodes, let's do it informally here by hand. To
claim your address space, simply add a 255 nodes block
below (please add consecutively and only claim 255 addresses at a time).

**Blocks**

NwkID 0:

01:FF:xx:xx reserved for testing

00:00:00:xx to 00:00:FF:xx reserved

00:01:00:xx to 00:01:FF:xx 1st 256 assigned blocks

00:02:00:xx to 00:02:FF:xx 2nd 256 blocks

etc.

## Registered Addresses
| Address block    |-| Who claimed it?        |
|------------------|-|------------------------|
| 01:FF:xx:xx      |-| *reserved for testing* |
| 00:01:00:xx      |-| name1 *(you?)*		  |
| 00:01:01:xx      |-| name2	   			  |
| 00:01:02:xx      |-| name3	   			  |
| etc.  		   |-| 						  |
|------------------|-|------------------------|
| DE:AD:BE:xx      | | turiphro               |

