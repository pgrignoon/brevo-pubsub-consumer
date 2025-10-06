Limitations actuelles:
    - toutes les tables sont dans le meme dataset
    - si on a plusieurs projets Brevo différents qui envoient le meme type de donnée (email transactionnel par ex) sur le target webhook, la donnée va se retrouver dans la meme table


solution: rajouter dans les headers de la requete envoyée coté Brevo les target dataset et tables pour flagger la donnée ?