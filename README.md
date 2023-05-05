# BLOCKCHAIN : RETRACE LA VIE DE TES PRODUITS

## Description

Blockchain permettant de retracer toute l'existence d'un produit. Chaque propriétaire peut transférer les droits de possession d'un objet à un autre utilisateur, qui valide la transaction avant de pouvoir le transférer à quelqu'un d'autre s'il le souhaite.

## Exemple

- 1er bloc : Tom crée cette blockchain liée à une banane qu'il vient de récolter
- 2nd bloc : Tom transmet la banane à Gérard le livreur
- 3ième bloc : Gérard le livreur valide la transaction
- 4ième bloc : Gérard vend la banane à Géraldine qui a très faim
- 5ième bloc : Géraldine valide la transaction
- 6ième bloc : au final Géraldine n'a pas faim... Elle transmet la banane à Jean-François

## Utilités

### Authenticité

Solution décentralisé (plus digne de confiance), qui permettra de visualiser tout le parcours d'un produit pour s'assurer de son origine.

### Qualité et confiance

Assurer un certain niveau de confiance envers certains produits en ajoutant l'accès à des contrôles qualités ou contrôles techniques dans la blockchain. Par exemple, en scannant le QR code liée à une voiture, vous pourriez consulter la date du dernier contrôle technique effectué, et l'identité du mécanicien s'en étant occupé.

### Sécurité

La personne ayant fermé le dernier block est la personne en possession de l'objet : très facile de repérer un vol.

## Etat du projet

Le projet est en cours de développement. Ci-dessous vous retrouverez le contenu des différents dossiers.

### User program

Contient un programme en go tournant uniquement en local permettant de créer votre utilisateur, créer une blockchain, ajouter des blocks à une blockchain. Pour employer le programme :

```bash
$ git clone https://github.com/geoffrey-diederichs/Blockchain
$ cd Blockchain/User_program
$ go run main.go
```

    Certaines fonctionnalités disponibles :

```bash
$ go run main.go -u # affiche votre id d'utilisateur

$ go run main.go -s # affiche les blockchains
$ go run main.go -s 1 # affiche la blocckchain d'id 1
$ go run main.go -s 1 1 # affiche le block 1 de la blockchain d'id 1

$ go run main.go -n # crée une blockchain
$ go run main.go -b 1 2 # ajoute un block à la blockchain d'id 1 transferant (ou récupérant) la propriété à l'utilisateur 2

$ go run main.go -v # vérifie la validité de toutes les blockchains
$ go run main.go -v 1 # vérifie la validité de la blockchain 1
```

### Site

Contient un prototype de site qui permettra dans le futur de consulter l'histoire de n'importe quel produit dans la blockchain depuis une page web ouverte à tous. Pour employer le site :

```bash
$ git clone https://github.com/geoffrey-diederichs/Blockchain
$ cd Blockchain/Site
$ go run main.go
```
Puis allez consulter les adresses localhost:8080/home?id=0 et localhost:8080/home?id=1.

### P2P

Code basique permettant un transfert de donné en P2P entre deux machines. Depuis l'hôte :

```bash
$ git clone https://github.com/geoffrey-diederichs/Blockchain
$ cd Blockchain/p2p/p2p_host
$ go run main.go
```

Depuis le client :

```bash
$ git clone https://github.com/geoffrey-diederichs/Blockchain
$ cd Blockchain/p2p/p2p_client
$ go run main.go
```
    
Un message apparaîtra sur l'hôte.