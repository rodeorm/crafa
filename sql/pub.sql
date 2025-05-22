ALTER SYSTEM SET wal_level = 'logical';

CREATE PUBLICATION users_pub FOR TABLE cmn.users;
CREATE PUBLICATION teams_pub FOR TABLE cmn.teams;
CREATE PUBLICATION user_teams_pub FOR TABLE cmn.userteams;

CREATE PUBLICATION areas_pub FOR TABLE ref.areas;
CREATE PUBLICATION categories_pub FOR TABLE ref.categories;
CREATE PUBLICATION iterations_teams_pub FOR TABLE ref.iterations;
CREATE PUBLICATION levels_pub FOR TABLE ref.levels;
CREATE PUBLICATION priorities_teams_pub FOR TABLE ref.priorities;
CREATE PUBLICATION statuses_teams_pub FOR TABLE ref.statuses;