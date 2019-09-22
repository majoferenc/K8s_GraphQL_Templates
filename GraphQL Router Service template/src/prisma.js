import { Prisma } from "prisma-binding";
import { fragmentReplacements } from "./resolvers/index";

let prisma_endpoint = "http://localhost:4466";

if (process.env.PRISMA_ENDPOINT !== undefined) {
  prisma_endpoint = process.env.PRISMA_ENDPOINT;
}

const prisma = new Prisma({
  typeDefs: "src/generated/prisma.graphql",
  endpoint: prisma_endpoint,
  fragmentReplacements
});

export { prisma as default };
