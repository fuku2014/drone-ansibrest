Use this plugin to execute ansible.

## Config

The following parameters are used to configure the plugin:

* **endpoint** - endpoint
* **playbook** - playbook
* **inventry** - inventry

## Example

Common example to execute ansible:

```yaml
pipeline:
  ansibrest:
    image: fuku2014/ansibrest
    endpoint: http://example.com:2400
	playbook: test.yml
	inventry: production
```
