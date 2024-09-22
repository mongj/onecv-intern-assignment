-- +migrate Up
CREATE TYPE SEX as ENUM ('male', 'female');
CREATE TYPE EMPLOYMENT_STATUS as ENUM ('employed', 'unemployed');
CREATE TYPE MARITAL_STATUS as ENUM ('single', 'married', 'widowed', 'divorced');
CREATE TYPE RELATION as ENUM ('parent', 'child', 'sibling', 'spouse', 'other');

CREATE TABLE people (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  sex SEX NOT NULL,
  date_of_birth DATE NOT NULL,
  employment_status EMPLOYMENT_STATUS NOT NULL,
  marital_status MARITAL_STATUS NOT NULL
);

CREATE TABLE households (
  id SERIAL PRIMARY KEY,
  person_id UUID NOT NULL,
  relative_id UUID NOT NULL,
  relation RELATION NOT NULL,
  FOREIGN KEY (person_id) REFERENCES people(id),
  FOREIGN KEY (relative_id) REFERENCES people(id)
);

CREATE TABLE applicants (
  person_id UUID PRIMARY KEY UNIQUE NOT NULL,
  FOREIGN KEY (person_id) REFERENCES people(id)
);

CREATE TABLE schemes (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL
);

CREATE TABLE scheme_benefits (
  id SERIAL PRIMARY KEY,
  scheme_id UUID NOT NULL,
  benefit_type SMALLINT NOT NULL,
  benefit_value VARCHAR(255) NOT NULL,
  benefit_description TEXT,
  FOREIGN KEY (scheme_id) REFERENCES schemes(id)
);

CREATE TABLE scheme_criteria (
  id SERIAL PRIMARY KEY,
  scheme_id UUID NOT NULL,
  criteria_type SMALLINT NOT NULL,
  criteria_value VARCHAR(255) NOT NULL,
  FOREIGN KEY (scheme_id) REFERENCES schemes(id)
);

CREATE TABLE applications (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  applicant_id UUID NOT NULL,
  scheme_id UUID NOT NULL,
  application_status SMALLINT NOT NULL,
  FOREIGN KEY (applicant_id) REFERENCES applicants(person_id),
  FOREIGN KEY (scheme_id) REFERENCES schemes(id)
);

-- +migrate Down
DROP TABLE applications;
DROP TABLE scheme_criteria;
DROP TABLE scheme_benefits;
DROP TABLE schemes;
DROP TABLE applicants;
DROP TABLE households;
DROP TABLE people;

DROP TYPE RELATION;
DROP TYPE MARITAL_STATUS;
DROP TYPE EMPLOYMENT_STATUS;
DROP TYPE SEX;