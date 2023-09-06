-- CreateTable
CREATE TABLE "users" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,
    "items" TEXT[],

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);
