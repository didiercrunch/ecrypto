
### public key
Note that any private key are a public key.


### private key



### cypher format

We use the `.cyp` file extension for the encrypted file.  A .cyp is nothing
else than a zip file containing the following shit.

1.  a `data.encrypted` containing the encrypted data
2.  a `meta_data.yml` file containing the meta data of the encrypted file
    *  symetric encryption algorithm
    *  mode
    *  encrypted symetric key
    *  hash of the encrypted data (always use full for validation)
3.  a 'signature' file with the signature of the `meta_data.yml` file
    func (this *RsaSha512) encrypt() ([]byte, error) {}


~~~

	 _____________________________________________________________________________
	|           
	|	Envelop:  zip file, no compression, no encryption
	|	 ________________________________________________________________________
	|	|	
	|	|	Paiload key:  cyphered by target pulic key
	|	|________________________________________________________________________
	|
	|	 _________________________________________________________________________
	|	|	
	|	|	Envelop meta data:  encryption type, mode, Paiload signature (yaml), IV,..
	|	|__________________________________________________________________________
	|
	|	 __________________________________________________________________________
	|	|	
	|	|	Paiload:  Zip file, no compression, symetrcally encrypted
	|	|	 _____________________________________________________________________
	|	|	|
	|	|	|	paiload intern metadata: Data signature, sender, message ... 
	|	|	|______________________________________________________________________
	|	|	
	|	|	 _______________________________________________________________________
	|	|	|
	|	|	|	Data:  Can be anything.  But for file and folders, it will be zipped
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|
	|	|	|____________________________________________________________________
	|	|
	|	|_________________________________________________________________________
	|
	|______________________________________________________________________________
	
~~~