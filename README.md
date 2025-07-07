# QIt

## Objetivos del Proyecto

- Crear una alternativa liviana a RabbitMQ o Kafka, con menos complejidad.
- Demostrar conocimientos avanzados en **Go**, **concurrencia con goroutines**, **canales**, y **APIs (HTTP/gRPC)**.
- Diseñar un sistema desacoplado, concurrente y extensible.
- Permitir a usuarios enviar, recibir, y reconocer (ack) mensajes de forma simple y eficiente.

---

## 🔄 Flujo de Trabajo de Qit

```
[ Producer ] → [ API HTTP/gRPC ] → [ Queue FIFO Concurrente ] → [ Consumer ]
```

1. **Producer**: Envía un mensaje a la cola.
2. **API**: Recibe el mensaje y lo encola.
3. **Queue**: Administra mensajes de forma concurrente.
4. **Consumer**: Recupera mensajes por HTTP o gRPC y los procesa.

---

## 🔹 Funcionalidades Clave

- Creación y eliminación de colas.
- Encolado y consumo de mensajes.
- Reconocimiento de mensajes (ack).
- Retransmisión en caso de timeout o falta de ack.
- Consumo por HTTP (polling/long-poll) y por gRPC (stream).
- Manejo concurrente con goroutines y canales.

---

## 📊 Estructura del Proyecto

```
qit/
├── cmd/
│   ├── server/        # Servidor HTTP y gRPC
│   └── cli/           # Cliente CLI opcional
├── internal/
│   ├── api/            # HTTP Handlers
│   ├── grpc/           # Implementación de gRPC
│   ├── queue/          # Estructuras de colas, buffer, workers
│   ├── broker/         # Administrador de colas y ruteo
│   └── storage/        # Persistencia opcional (memoria, archivo)
├── proto/              # Archivos .proto
└── pkg/                # Utilidades, logging, configuración
```

---

## 📅 Uso Típico (HTTP)

### 1. Crear una cola

```
curl -X POST http://localhost:8080/queues/email-jobs
```

### 2. Enviar un mensaje

```
curl -X POST http://localhost:8080/queues/email-jobs/messages \
     -H "Content-Type: application/json" \
     -d '{"to": "user@example.com", "subject": "Hello!"}'
```

### 3. Recibir un mensaje

```
curl http://localhost:8080/queues/email-jobs/messages
```

### 4. Reconocer el mensaje (ACK)

```
curl -X DELETE http://localhost:8080/queues/email-jobs/messages/msg-1234
```

---

## 🔗 Uso Típico (gRPC)

- `SendMessage()` para encolar.
- `ReceiveMessages()` como stream bidireccional.
- `AckMessage()` para reconocimiento.
