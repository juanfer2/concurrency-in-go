# The sync package

Este paquete contiene primitives concurrency, que son más útiles para la sincronización a la memoria
de bajo nivel, varios idiomas lo tienen la diferencia es que go ha construido unas nuevas primitives concurrency en la parte superior de la sincronización proporciona una amplia herramientas que nos pueden ayudar.estas operaciones tienen su uso, principalmente en ámbitos pequeños, como una estructura. Depende de ti decidir cuando la sincronización del acceso a la memoria es apropiada.


## WaitGroup

Es una excelente manera de esperar a que se complente un conjunto de operaciones simultáneas 
cuando no le importa el resultado de la operación simultánea, o tiene otros medios de recopilar sus resultados. Si ninguna de esas condiciones es cierta, le sugiero que use channels y una declaración de selección en su lugar.
