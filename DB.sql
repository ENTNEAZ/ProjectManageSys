CREATE TABLE research_room (
    research_room_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    research_room_name VARCHAR(255) NOT NULL,
    research_room_direction VARCHAR(255) NOT NULL
);

CREATE TABLE working_area (
    working_area_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    working_area_size INT NOT NULL,
    working_area_address VARCHAR(255) NOT NULL
);

CREATE TABLE research_room_working_area (
    research_room_id INT NOT NULL,
    working_area_id INT NOT NULL,
    PRIMARY KEY (research_room_id, working_area_id),
    FOREIGN KEY (research_room_id) REFERENCES research_room (research_room_id),
    FOREIGN KEY (working_area_id) REFERENCES working_area (working_area_id),
    CHECK (research_room_id IN (SELECT research_room_id FROM research_room)),
    CHECK (working_area_id IN (SELECT working_area_id FROM working_area)),
    CHECK NOT EXISTS (SELECT * FROM research_room_working_area AS r WHERE working_area_id = r.working_area_id)
);

CREATE TABLE worker(
    worker_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    worker_name VARCHAR(255) NOT NULL,
    worker_gender VARCHAR(255) NOT NULL,
    worker_birth DATE NOT NULL,
    worker_join_time DATE NOT NULL,
    worker_job VARCHAR(255) NOT NULL
);

CREATE TABLE research_room_boss (
    research_room_id INT NOT NULL,
    worker_id INT NOT NULL,
    join_time DATE NOT NULL,
    term INT NOT NULL,
    PRIMARY KEY (research_room_id, worker_id),
    FOREIGN KEY (research_room_id) REFERENCES research_room (research_room_id),
    FOREIGN KEY (worker_id) REFERENCES worker (worker_id)
);

CREATE TABLE research_room_sectary (
    research_room_id INT NOT NULL,
    worker_id INT NOT NULL,
    job_detail VARCHAR(255) NOT NULL,
    PRIMARY KEY (research_room_id, worker_id),
    FOREIGN KEY (research_room_id) REFERENCES research_room (research_room_id),
    FOREIGN KEY (worker_id) REFERENCES worker (worker_id)
);

CREATE TABLE research_room_worker (
    research_room_id INT NOT NULL,
    worker_id INT NOT NULL,
    direction VARCHAR(255) NOT NULL,
    PRIMARY KEY (research_room_id, worker_id),
    FOREIGN KEY (research_room_id) REFERENCES research_room (research_room_id),
    FOREIGN KEY (worker_id) REFERENCES worker (worker_id)
);

CREATE TABLE project (
    project_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    worker_id INT NOT NULL,
    project_name VARCHAR(255) NOT NULL,
    project_detail VARCHAR(255) NOT NULL,
    project_start_time DATE NOT NULL,
    project_end_time DATE NOT NULL,
    project_fund INT NOT NULL,
    project_participant_id INT NOT NULL,
    FOREIGN KEY (worker_id) REFERENCES worker (worker_id),
    FOREIGN KEY (project_participant_id) REFERENCES project_participant (project_participant_id)
);

CREATE TABLE project_participant (
    project_participant_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    project_participant_name VARCHAR(255) NOT NULL,
    project_participant_address VARCHAR(255) NOT NULL,
    project_participant_worker_id INT NOT NULL,
    FOREIGN KEY (project_participant_worker_id) REFERENCES project_participant_worker (project_participant_worker_id)
);

CREATE TABLE project_participant_worker (
    project_participant_worker_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    project_participant_worker_telephone VARCHAR(255) NOT NULL,
    project_participant_worker_mobile VARCHAR(255) NOT NULL,
    project_participant_worker_email VARCHAR(255) NOT NULL
);

CREATE TABLE project_participant_project_participant_worker_contact (
    project_participant_id INT NOT NULL,
    project_participant_worker_id INT NOT NULL,
    PRIMARY KEY (project_participant_id, project_participant_worker_id),
    FOREIGN KEY (project_participant_id) REFERENCES project_participant (project_participant_id),
    FOREIGN KEY (project_participant_worker_id) REFERENCES project_participant_worker (project_participant_worker_id)
);

CREATE TABLE project_worker (
    project_id INT NOT NULL,
    worker_id INT NOT NULL,
    PRIMARY KEY (project_id, worker_id),
    FOREIGN KEY (project_id) REFERENCES project (project_id),
    FOREIGN KEY (worker_id) REFERENCES worker (worker_id)
);

CREATE TABLE sub_project (
    sub_project_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    project_id INT NOT NULL,
    worker_id INT NOT NULL,
    sub_project_end_time DATE NOT NULL,
    sub_project_fund INT NOT NULL,
    sub_project_tech_detail VARCHAR(255) NOT NULL,
    FOREIGN KEY (project_id) REFERENCES project (project_id),
    FOREIGN KEY (worker_id) REFERENCES worker (worker_id)
);

CREATE TABLE sub_project_worker (
    sub_project_id INT NOT NULL,
    worker_id INT NOT NULL,
    join_time DATE NOT NULL,
    sub_project_worker_fund INT NOT NULL,
    workload INT NOT NULL,
    PRIMARY KEY (sub_project_id, worker_id),
    FOREIGN KEY (sub_project_id) REFERENCES sub_project (sub_project_id),
    FOREIGN KEY (worker_id) REFERENCES worker (worker_id)
);

CREATE TABLE project_project_fruit (
    project_id INT NOT NULL,
    project_fruit_id INT NOT NULL,
    project_fruit_name VARCHAR(255) NOT NULL,
    project_fruit_get_time DATE NOT NULL,
    worker_id INT NOT NULL,
    project_fruit_master_rank INT NOT NULL,
    PRIMARY KEY (project_id, project_fruit_id),
    FOREIGN KEY (project_id) REFERENCES project (project_id),
    FOREIGN KEY (project_fruit_id) REFERENCES project_fruit (project_fruit_id),
    FOREIGN KEY (worker_id) REFERENCES worker (worker_id)
);

CREATE TABLE project_fruit (
    project_fruit_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    project_fruit_type VARCHAR(255) NOT NULL,
    project_fruit_detail VARCHAR(255) NOT NULL
);