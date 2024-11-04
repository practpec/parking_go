# Estructura de datos y configuración

```go
estructura Parking {
    espacios: canal de enteros
    mutexEntrada: mutex
    espaciosArray: array de booleanos // Indica si un espacio está ocupado
}

// Inicializar el estacionamiento
función NuevaParking(canalEspacios, mutexEntrada) {
    crear instancia de Parking
    retornar instancia
}
```

## Clase para coches

```go
estructura Car {
    id: entero
    tiempoEstacionamiento: duración
}

// Inicializar coche
función NuevoCar(id) {
    crear imagen de coche aleatoria
    asignar tiempo de estacionamiento aleatorio
    retornar nueva instancia de Car
}
```

## Función principal del sistema

```go
función Main() {
    // Crear parking y otros recursos
    parking = NuevaParking(canal de enteros, mutex)
    wg = nueva WaitGroup()

    // Generar coches
    para i desde 0 hasta N {
        wg.Agregar(1)
        id = i
        lanzar goroutine para crear coche(id)
    }
    wg.Esperar()
}
```

## Función para manejar la entrada del coche

```go
función EntrarCarro(coche, parking) {
    // Enviar ID del coche al canal de espacios
    parking.espacios <- coche.id

    // Bloquear el acceso al estacionamiento
    parking.mutexEntrada.Lock()
    
    // Esperar un corto periodo para simular entrada
    esperar 300 ms
    
    // Obtener el estado de los espacios
    espaciosDisponibles = parking.espaciosArray
    si todos los espacios están ocupados {
        desbloquear el mutex
        retornar // No hay espacios disponibles
    }

    // Encontrar un espacio disponible
    espacioIndex = seleccionarAleatoriamente(espaciosDisponibles)
    espaciosDisponibles[espacioIndex] = verdadero // Marcar espacio como ocupado
    coche.espacio = espacioIndex // Asignar espacio al coche

    // Actualizar el estado del parking
    parking.espaciosArray = espaciosDisponibles
    parking.mutexEntrada.Unlock()
}
```

## Función para manejar la salida del coche

```go
función SalirCarro(coche, parking) {
    // Bloquear el acceso al estacionamiento
    parking.mutexEntrada.Lock()

    // Retirar ID del canal de espacios
    <-parking.espacios

    // Actualizar estado de espacios
    espaciosDisponibles = parking.espaciosArray
    espaciosDisponibles[coche.espacio] = falso // Marcar espacio como libre
    parking.espaciosArray = espaciosDisponibles

    // Desbloquear el mutex
    parking.mutexEntrada.Unlock()
}
```

## Función que simula el aparcamiento de un coche

```go
función Aparcar(coche, parking, wg) {
    coche.Entrar(parking)
    
    // Esperar el tiempo de aparcamiento
    esperar coche.tiempoEstacionamiento
    
    coche.Salir(parking)
    wg.Done() // Indicar que este goroutine ha terminado
}
```

## Lanzar la creación de coches

```go
función crearCoche(id) {
    coche = NuevoCar(id)
    Aparcar(coche, parking, wg) // Iniciar el proceso de aparcamiento
}
```