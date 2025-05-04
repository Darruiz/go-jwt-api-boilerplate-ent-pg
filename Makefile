# Cria um novo schema do Ent para uma entidade dentro do padrão Vertical Slice
# Exemplo de uso:
#   make ent-new name=User
# Vai gerar: internal/features/user/ent/schema/user.go
ent-new:
	ent new $(name) --target internal/features/$(shell echo $(name) | tr A-Z a-z)/ent/schema

# Gera o código (client) do Ent com base no schema da feature especificada
# Exemplo de uso:
#   make ent-generate name=User
# Vai gerar o client dentro de: internal/features/user/ent/
ent-generate:
	go run entgo.io/ent/cmd/ent generate internal/features/$(shell echo $(name) | tr A-Z a-z)/ent/schema

# Gera todos os clients do Ent para todas as features existentes
# Útil após clonar o projeto ou quando modificar múltiplos schemas
# Exemplo de uso:
#   make ent-generate-all
ent-generate-all:
	find internal/features -type d -name schema | while read dir; do \
		go run entgo.io/ent/cmd/ent generate $$dir; \
	done


# Aplica as migrations com base no schema Ent + Atlas
ent-migrate:
	docker run --rm \
		-v $(PWD):/app \
		-w /app \
		arigaio/atlas \
		schema apply \
		--url "postgres://postgres:postgres@db:5432/dzfinance?sslmode=disable" \
		--dev-url "sqlite://file?mode=memory&cache=shared&_fk=1" \
		--to "file://internal/migrations" \
		--auto-approve
# Roda a aplicação localmente
# Exemplo de uso:
#   make dev
dev:
	go run cmd/main.go





### 🐳 DOCKER: tudo de atalho docker aqui pra baixo
### ─────────────────────────────

# Builda os containers com cache (rápido)
docker-build:
	docker-compose build

# Builda os containers sem cache (seguro para atualizações de dependências)
docker-build-nc:
	docker-compose build --no-cache

# Sobe a stack com rebuild inteligente
docker-up:
	docker-compose up -d --build

# Derruba tudo (inclusive volumes, redes e cache)
docker-down:
	docker-compose down -v --remove-orphans

# Força rebuild + up limpo
docker-restart:
	make docker-down
	make docker-up

# Acompanha logs de todos os containers (CTRL+C para sair)
docker-logs:
	docker-compose logs -f --tail=100

# Acessa o container da API com shell interativo
docker-api-sh:
	docker exec -it go-api sh

# Acessa o container do Postgres com psql
docker-db:
	docker exec -it go-db psql -U postgres -d dzfinance