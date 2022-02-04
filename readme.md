#### Developing a RESTful API with Go and Gin

###### This is an example from [Go documentation](https://go.dev/doc/tutorial/web-service-gin) with some notes.


1) on Windows 10, *curl* is an alias for the 'Invoke-WebRequest' PowerShell *cmdlet*, which has its own parameter passing:

	SYNTAX
	
	
	Invoke-WebRequest [-Uri] <Uri> [-Body <Object>] [-Certificate <X509Certificate>] [-CertificateThumbprint <String>] [-ContentType <String>] [-Credential <PSCredential>] [-DisableKeepAlive] [-Headers <IDictionary>] [-InFile <String>] [-MaximumRedirection <Int32>] [-Method {Default | Get | Head | Post | Put | Delete | Trace | Options | Merge | Patch}] [-OutFile <String>] [-PassThru] [-Proxy <Uri>] [-ProxyCredential <PSCredential>] [-ProxyUseDefaultCredentials] [-SessionVariable <String>] [-TimeoutSec <Int32>] [-TransferEncoding {chunked | compress | deflate | gzip | identity}] [-UseBasicParsing] [-UseDefaultCredentials] [-UserAgent <String>] [-WebSession <WebRequestSession>] [<CommonParameters>]

	
Thus, *curl* should be executed in a Windows environment like this (relative to this code):

	curl http://localhost:8080/albums -ContentType "application/json" -Method Post -Body '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'



2) for some reason in Ubuntu 20 (LTS) I was consistently getting "404 Not Found" in response to the request 'curl http://localhost:8080/albums/2'
   if you know why, let me know.
	
