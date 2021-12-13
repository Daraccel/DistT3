* -> por hacer
- -> realizado

Servidores falcrum
    - tienen los 3 los archivos
        archivos:
            - Formato: Nombre_Planeta Nombre_Ciudad Cantidad_Soldados_Rebeldes
            - almacenar segun planeta

    - cada servidor posee [x,y,z] por archivos
    
    *Log de registro con modificacion de registros
        - Formato: accion planeta_afectado ciudad_afectada Nuevo_Valor
        * borrar y crear nuevo vacio 
    
    * Mantener consistencia eventual cada 2 min
    
    * merge si hay diferencias en [x,y,z] de forma automatica
    
    // pensamientos multithreading para recibir llamados a accciones y actualizar archivos con seccion protegida [x,y,z]?, por probar como se comporta con un servidor reciviendo acciones al mismo planeta

* Broker Mos Eisley
    - redirige informantes a replica aleatoria
    
    * caso princesa leia, se comporta como intermediario (* Revizar bien a que se refiere este punto)
    
    * caso de conflicto de informacion -> mandar a replica en especifico

* Informantes
    * acciones
        * AddCity nombre_planeta nombre_ciudad nuevo_valor
            - No existe planeta -> se crea nuevo registro de planeta
            * en caso de que no haya valor -> valor = 0
        
        - UpdateName nombre_planeta nombre_ciudad nuevo_valor
            update nomber ciudad en planeta indicado
        
        - DeleteCity nombre_planeta nombre_ciudad
            eliminar ciudad en registro planetario indicado
        
    - Comunicacion
        1- mandan a broker mos eisley comando
        2- broker responde con direccion de fulcrum a mandar
        3- re-enviar mismo mensaje enviado a broker hacia fulcrum
        4- recivir reloj vector de archivo modificado
        
    - modelo consistencia read your writes
        con cada write, hay que hacer read al objeto 
    
    - mantener en memoria informacion de los registros planetarios modificados con su reloj vector y direccion servidor
        creo que su buen map sirve, por
        
* Leia Organa
    * consultas a broker para saber cantidad de rebeldes por ciudad en planeta
        GetNumberRebeldes nombre_planeta nombre_ciudad
            recive de vuelta numero de rebeldes en ciudad de ese planeta y reloj vectorial de registro de planeta
    
    * consistencia monotonic reads
        debe completarse el write antes de que el mismo proceso haga otro write en el mismo data item
    
    * informacion en memoria de ciudades que ha solicitado, reloj vectorial y servidor

* Consistencia
    - tener relojes vectores [x,y,z] en cada servidor por archivo
        como guardarlos por archivo?
    
    * merge
        * definir nodo dominante
            * nodo dominante es base para aplicar cambios en otros nodos

        * Utilizar log de cambios de servidores para actualizar registro segun reloj vectores

        * dejar registrado como lidiamos conflictos particulares
Otras notas:

generar .proto
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
funciones.proto

problema:
protoc-gen-go: program not found or is not executable
solucion:
sudo apt install protobuf-compiler
sudo apt install golang-goprotobuf-dev
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH=$PATH:$(go env GOPATH)/bin
