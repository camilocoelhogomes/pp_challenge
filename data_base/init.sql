CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE SEQUENCE "public".user_internal_id_seq START WITH 1 INCREMENT BY 1;
CREATE TABLE "public".document_type (
  id smallserial NOT NULL,
  name varchar(100),
  CONSTRAINT pk_document_type PRIMARY KEY (id),
  CONSTRAINT unq_document_type UNIQUE (name)
);
CREATE TABLE "public"."user" (
  internal_id bigserial NOT NULL,
  id uuid DEFAULT uuid_generate_v4(),
  name varchar(255) NOT NULL,
  document_type smallint NOT NULL,
  email varchar(255) NOT NULL,
  document_number varchar(255) NOT NULL,
  CONSTRAINT user_email_key UNIQUE (email),
  CONSTRAINT user_document_number_key UNIQUE (document_number),
  CONSTRAINT unq_user_internal_id UNIQUE (internal_id),
  CONSTRAINT pk_user PRIMARY KEY (internal_id),
  CONSTRAINT unq_user_document_type UNIQUE (document_type),
  CONSTRAINT fk_user_document_type FOREIGN KEY (document_type) REFERENCES "public".document_type(id)
);
CREATE TABLE "public".authentication (
  internal_id bigserial NOT NULL,
  credential varchar(255),
  CONSTRAINT pk_authentication PRIMARY KEY (internal_id),
  CONSTRAINT fk_authentication_user FOREIGN KEY (internal_id) REFERENCES "public"."user"(internal_id)
);
INSERT INTO "public".document_type(id, name)
VALUES (1, 'cpf');
INSERT INTO "public".document_type(id, name)
VALUES (2, 'cnpj');