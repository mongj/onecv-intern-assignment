DO $$
-- Scheme
DECLARE s1 CONSTANT UUID := gen_random_uuid();
DECLARE s2 CONSTANT UUID := gen_random_uuid();

-- People
DECLARE p1 CONSTANT UUID := gen_random_uuid();
DECLARE p2 CONSTANT UUID := gen_random_uuid();
DECLARE c1 CONSTANT UUID := gen_random_uuid();
DECLARE c2 CONSTANT UUID := gen_random_uuid();

BEGIN
  -- Clear existing data
	DELETE FROM scheme_benefits;
	DELETE FROM scheme_criteria;
	DELETE FROM schemes;

  DELETE FROM applicants;
  DELETE FROM households;
  DELETE FROM people;

  -- Seed schemes data
	INSERT INTO schemes (id, name) VALUES (s1, 'Retrenchment Assistance Scheme');
	INSERT INTO schemes (id, name) VALUES (s2, 'Retrenchment Assistance Scheme (families)');

	INSERT INTO scheme_criteria (scheme_id, criteria_key, criteria_value) VALUES (s1, 0, 'unemployed');
	INSERT INTO scheme_criteria (scheme_id, criteria_key, criteria_value) VALUES (s2, 0, 'unemployed');
	INSERT INTO scheme_criteria (scheme_id, criteria_key, criteria_value) VALUES (s2, 2, 'true');
	INSERT INTO scheme_criteria (scheme_id, criteria_key, criteria_value) VALUES (s2, 3, 'primary');

	INSERT INTO scheme_benefits (scheme_id, description, amount) VALUES (s1, 'SkillsFuture Credits', 500.00);
	INSERT INTO scheme_benefits (scheme_id, description, amount) VALUES (s2, 'SkillsFuture Credits', 500.00);
	INSERT INTO scheme_benefits (scheme_id, description, amount) VALUES (s2, 'Daily school meal vouchers', 8.00);

  -- Seed people data
  INSERT INTO people (id, name, sex, date_of_birth, employment_status, marital_status, current_school_level) 
    VALUES (p1, 'John Doe', 'male', '1990-01-01', 'unemployed', 'single', NULL);
  INSERT INTO people (id, name, sex, date_of_birth, employment_status, marital_status, current_school_level)
    VALUES (p2, 'Jane Doe', 'female', '1990-01-01', 'unemployed', 'married', NULL);
  INSERT INTO people (id, name, sex, date_of_birth, employment_status, marital_status, current_school_level) 
    VALUES (c1, 'Child 1', 'male', '2018-01-01', 'unemployed', 'single', 'preschool');
  INSERT INTO people (id, name, sex, date_of_birth, employment_status, marital_status, current_school_level) 
    VALUES (c2, 'Child 2', 'female', '2016-01-01', 'unemployed', 'single', 'primary');
  
  INSERT INTO households (person_id, relative_id, relation) VALUES (p1, c1, 'child');
  INSERT INTO households (person_id, relative_id, relation) VALUES (p2, c2, 'child');

  INSERT INTO applicants (person_id) VALUES (p1);
  INSERT INTO applicants (person_id) VALUES (p2);
END $$;
