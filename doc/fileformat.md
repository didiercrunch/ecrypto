
### public contract

A public contract is the information a user publish to allow anybody to send 
him encrypted files.  Of course the public contract contains the user
public key but it contains much more.

The contract itself does not need to be sign.  The signature is level of
complexity above the public contract


~~~
---
#  example of a  contract

name: john dear

organizarion: JohnDeer corporation

email: johndear@johndeer.co

url: "the url where to fiofbnd an exact copy of the contract"

root_url: "the url where to find the most up to date version of the contract"

accepted_hashs:
	- sha512

accepted_block_cypher:
	- aes

accepted_modes:
	- ofb

signatures_required: yes

rsa_public_keys:
	-
		n: "hexadecimal n"
		e: 3424
~~~


### public key
Note that any private key are a public key.


### private key



### cypher format

We use the `.filou` file extension for the encrypted file.  A .filou is nothing
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