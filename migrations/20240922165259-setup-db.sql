-- +migrate Up

CREATE TABLE people (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  sex SMALLINT NOT NULL,
  dob DATE NOT NULL,
  employment_status SMALLINT NOT NULL,
  marital_status SMALLINT NOT NULL
);

CREATE TABLE relationships (
  id SERIAL PRIMARY KEY,
  person_id UUID NOT NULL,
  relative_id UUID NOT NULL,
  relationship_type SMALLINT NOT NULL,
  FOREIGN KEY (person_id) REFERENCES people(id),
  FOREIGN KEY (relative_id) REFERENCES people(id)
);

CREATE TABLE applicants (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  person_id UUID NOT NULL,
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
  FOREIGN KEY (applicant_id) REFERENCES applicants(id),
  FOREIGN KEY (scheme_id) REFERENCES schemes(id)
);

-- +migrate Down
DROP TABLE applications;
DROP TABLE scheme_criteria;
DROP TABLE scheme_benefits;
DROP TABLE schemes;
DROP TABLE applicants;
DROP TABLE relationships;
DROP TABLE people;