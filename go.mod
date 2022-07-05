module app

go 1.18

replace gitlab.com/jonas.jasas/httprelay => gitlab.com/shynome/httprelay v0.1.0

require (
	github.com/donovanhide/eventsource v0.0.0-20210830082556-c59027999da0
	gitlab.com/jonas.jasas/httprelay v0.0.0-20220331143220-3f804d44d5eb
)

require (
	gitlab.com/jonas.jasas/buffreader v0.0.0-20200406102452-cad14a5681d3 // indirect
	gitlab.com/jonas.jasas/bufftee v0.0.0-20190309090616-0f18f95bd4d2 // indirect
	gitlab.com/jonas.jasas/closechan v0.0.0-20200419060521-9f8eba9e316d // indirect
)
