# Azure-SQLDB

# Create SQL Server and Database

Prerequisite : The [Azure Service operator](https://github.com/Azure/azure-service-operator) Deployed, configured and running 

1- Create ResourceGroup, AzureSqlServer and AzureSqlDatabase.

 ```oc create -f manifests/azureSqlDb.yaml```

make sure all objects created successfully and by verifying from azure portal. 

## Verify connection locally 

1- update the main.go file add username , password and sql server connection setting

2- run `go run main.go` 

3- verify connection made with created database if you are facing firewall issue
add ip address in azure portal network settings or create the firewall object  `oc create -f manifests/azuresqlfirewallrule.yaml`


## Deploy example on kubernetes 

1- Create project/namespace 

2- update azureexample-deploy.yaml file secret and deployment objects

3- Deploy the application `oc create -f manifests/azureSqlDb.yaml`

4- check the link for more details presentation: https://drive.google.com/file/d/18BxZqkL8Ja4audNgBJX5ZGO7apt2TdZZ/view