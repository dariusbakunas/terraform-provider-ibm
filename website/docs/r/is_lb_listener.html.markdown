---

subcategory: "VPC infrastructure"
page_title: "IBM : lb_listener"
description: |-
  Manages IBM load balancer listener.
---

# ibm_is_lb_listener
Create, update, or delete a listener for a VPC load balancer. For more information, about load balancer listener, see [working with listeners](https://cloud.ibm.com/docs/vpc?topic=vpc-nlb-listeners).

**Note**

When provisioning the load balancer listener along with load balancer pool or pool member, Use explicit depends on the resources or perform the terraform apply with parallelism. For more information, about explicit dependencies, see [create resource dependencies](https://learn.hashicorp.com/terraform/getting-started/dependencies#implicit-and-explicit-dependencies).

## Example usage
An example, to create a load balancer listener along with the pool and pool member.

```terraform
resource "ibm_is_lb_listener" "testacc_lb_listener" {
  lb       = "8898e627-f61f-4ac8-be85-9db9d8bfd345"
  port     = "9080"
  protocol = "http"
  https_redirect_listener="r134-8c58bfe1-db02-4790-95ce-fe5bb892d78f"
  https_redirect_status_code=301
  https_redirect_uri="/example?doc=get"
}

resource "ibm_is_lb_pool" "webapptier-lb-pool" {
  lb                 = "8898e627-f61f-4ac8-be85-9db9d8bfd345"
  name               = "a-webapptier-lb-pool"
  protocol           = "http"
  algorithm          = "round_robin"
  health_delay       = "5"
  health_retries     = "2"
  health_timeout     = "2"
  health_type        = "http"
  health_monitor_url = "/"
  depends_on         = [ibm_is_lb_listener.testacc_lb_listener]
}

resource "ibm_is_lb_pool_member" "webapptier-lb-pool-member-zone1" {
  count          = "2"
  lb             = "8898e627-f61f-4ac8-be85-9db9d8bfd345"
  pool           = element(split("/", ibm_is_lb_pool.webapptier-lb-pool.id), 1)
  port           = "80"
  target_address = "192.168.0.1"
  depends_on     = [ibm_is_lb_listener.testacc_lb_listener]
}
```

### Sample to create a load balancer listener policy for a `https_redirect` action.

```terraform
resource "ibm_is_lb" "lb2"{
  name    = "mylb"
  subnets = ["35860fed-c911-4936-8c94-f0d8577dbe5b"]
}

resource "ibm_is_lb_listener" "lb_listener1"{
  lb       = ibm_is_lb.lb2.id
  port     = "9086"
  protocol = "https"
  certificate_instance="crn:v1:staging:public:cloudcerts:us-south:a2d1bace7b46e4815a81e52c6ffeba5cf:af925157-b125-4db2-b642-adacb8b9c7f5:certificate:c81627a1bf6f766379cc4b98fd2a44ed"
}

resource "ibm_is_lb_listener" "lb_listener2"{
  lb       = ibm_is_lb.lb2.id
  port     = "9087"
  protocol = "http"
  https_redirect_listener = ibm_is_lb_listener.lb_listener1.listener_id
  https_redirect_status_code = 301
  https_redirect_uri = "/example?doc=geta" 
}
```

## Timeouts
The `ibm_is_lb_listener` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating Instance.
- **update** - (Default 10 minutes) Used for updating Instance.
- **delete** - (Default 10 minutes) Used for deleting Instance.


## Argument reference
Review the argument references that you can specify for your resource. 

- `accept_proxy_protocol`- (Optional, Bool)  If set to **true**, listener forwards proxy protocol information that are supported by load balancers in the application family. Default value is **false**.
- `lb` - (Required, Forces new resource, String) The load balancer unique identifier.
- `port`- (Required, Integer) The listener port number. Valid range 1 to 65535.
- `protocol` - (Required, String) The listener protocol. Enumeration type are `http`, `tcp`, and `https`. Network load balancer supports only `tcp` protocol.
- `default_pool` - (Optional, String) The load balancer pool unique identifier.
- `certificate_instance` - (Optional, String) The CRN of the certificate instance, it is applicable(mandatory) only to https protocol.
- `connection_limit` - (Optional, Integer) The connection limit of the listener. Valid range is **1 to 15000**. Network load balancer do not support `connection_limit` argument.
- `https_redirect_listener` - (Optional, String) ID of the listener that will be set as http redirect target.
- `https_redirect_status_code` - (Optional, Integer) The HTTP status code to be returned in the redirect response, one of [301, 302, 303, 307, 308].
- `https_redirect_uri` - (Optional, String) Target URI where traffic will be redirected.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the load balancer listener.
- `status` - (String) The status of load balancer listener.

## Import
The `ibm_is_lb_listener` resource can be imported by using the load balancer ID and listener ID.

**Syntax**

```
$ terraform import ibm_is_lb_listener.example <loadbalancer_ID>/<listener_ID>
```

**Example**

```
$ terraform import ibm_is_lb_listener.example d7bec597-4726-451f-8a63-e61212c32c/cea6651a-bc0a-4438-9f8a-44444f3ebb
```
