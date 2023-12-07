
# ZincSearch

Se creo un contenedor en docker para correr zincsearch, ademas se crea un volumen para almacenar los datos. Creará una carpeta "data" dentro del directorio en el cual se corra el siguiente comando:

docker run -v %cd%/data:/data -e ZINC_DATA_PATH="/data" -p 4080:4080 -e ZINC_FIRST_ADMIN_USER=admin -e ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123 --name zincsearch public.ecr.aws/zinclabs/zincsearch:latest


# Rest API (GO)

El api es muy simple y se basa en que el usuario podrá realizar busquedas utilizando tres parámetros: emisor del emial, receptor del email, y titulo del email, los tres son opcionales. Zincsearch hara el match de los correos que cumplan con los filtros escritos. 


Por hacer:

- Recordar cambiar el url ".../test_emails/" por el nombre final del indice

# UI (Vue.js)

