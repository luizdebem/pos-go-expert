### SERVER

- **Só 1 endpoint:** `GET localhost:8080/cotacao`

- **Processamento:**
  - Consumir a API de câmbio entre Dólar e Real:
    - URL: `https://economia.awesomeapi.com.br/json/last/USD-BRL`
  - Retornar o resultado para o cliente no formato JSON.
  - Salvar a cotação obtida no SQLite.

- **Timeouts via context:**
  - Timeout da requisição da API de câmbio (Dólar-Real): 200ms.
  - Timeout para persistir no SQLite: 10ms.
  - Timeout deve retornar erro em log

---

### CLIENT
  - Consumir a endpoint do servidor.
  - Timeout máximo para a resposta do servidor: 300ms.
  - Salvar o resultado em um arquivo `cotacao.txt` no formato `"Dólar: {valor}\n"`
  - Timeout deve retornar erro em log
