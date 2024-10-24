Projet Final - Système de Gestion des Employés

Description
Ce projet a pour but de créer un système de gestion des employés pour une entreprise, permettant de stocker, afficher et manipuler les informations des employés, départements et postes. L'application utilise une base de données SQL pour le stockage des données, et une interface web conviviale pour interagir avec ces données.

Objectifs
Base de données SQL : Créer une base de données qui stocke les informations sur les employés, les départements et les postes.
Requêtes SQL : Utiliser des requêtes SQL pour extraire et afficher les informations sur une page web.
Interface Utilisateur : Développer une interface utilisateur en utilisant HTML et CSS pour une navigation et une expérience utilisateur optimales.
Formulaires HTML : Intégrer des formulaires pour permettre la saisie d'informations et leur envoi à la base de données.
Navigation : Créer des pages de navigation pour permettre aux utilisateurs de parcourir les différentes sections du site.
JavaScript : Utiliser JavaScript pour ajouter des fonctionnalités interactives à l'interface.
Technologies Utilisées
Langage de Programmation : Go (Golang)
Base de Données : SQLite
Web : HTML, CSS, JavaScript
Framework : Aucune dépendance externe pour le backend (utilisation de la bibliothèque standard de Go)
Installation
Cloner le dépôt :

bash
git clone https://github.com/Dupasronan/TP-SQL-FINAL.git
cd TP-SQL-FINAL

Installer Go : Assurez-vous que Go est installé sur votre machine. Vous pouvez le télécharger à partir de golang.org.

Installer les dépendances : Utilisez le gestionnaire de paquets de Go pour installer les dépendances :

bash
go get -u github.com/mattn/go-sqlite3

Lancer le serveur : Exécutez le fichier main.go pour démarrer le serveur :

bash
go run main.go

Accéder à l'application : Ouvrez votre navigateur et allez à l'adresse http://localhost:8080.

Fonctionnalités
Affichage des Employés : Une page qui liste tous les employés avec la possibilité de rechercher par prénom ou nom.
Ajout d'un Employé : Un formulaire pour ajouter un nouvel employé à la base de données.
Suppression d'un Employé : Une fonctionnalité pour supprimer un employé en fonction de son identifiant.

Structure des Dossiers

bash
Copier le code
/<>
│
├── employees.css       # style pour employees.html
├── employees.html      # Page affichant la liste des employés
├── Entreprise.db       # Base de données SQLite
├── gestion.css         # style pour gestion.html
├── gestion.html        # Page de gestion des employés
├── index.html          # Page d'accueil
├── main.go             # Fichier principal de l'application
└── style.css           # style pour style.html

Auteurs
Ondo Ella Fils Noé - Dupas Ronan