#################################
# STEP 1 ~ Install dependencies #
#################################
FROM oven/bun:1.2.2-alpine AS deps

WORKDIR /app

# Install dependencies
COPY package.json bun.lock ./
RUN bun install --frozen-lockfile


####################################
# STEP 2 ~ Create production build #
####################################
FROM oven/bun:1.2.2-alpine AS builder
WORKDIR /app

# Next.js collects completely anonymous telemetry data about general usage.
# Learn more here: https://nextjs.org/telemetry
# Uncomment the following line in case you want to disable telemetry during the build.
ENV NEXT_TELEMETRY_DISABLED=1
ENV BUILD_STANDALONE=true

# Copy dependencies from step 1
COPY --from=deps /app/node_modules ./node_modules
COPY . .

RUN bun run build

#################################
# STEP 3 ~ Run production build #
#################################
FROM oven/bun:1.2.2-alpine AS runner
WORKDIR /app

EXPOSE 3000
ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1

# Copy public assets
COPY --from=builder /app/public ./public

# Set the correct permission for prerender cache
RUN mkdir .next

# Automatically leverage output traces to reduce image size
# https://nextjs.org/docs/advanced-features/output-file-tracing
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static

CMD ["bun", "server.js"]
