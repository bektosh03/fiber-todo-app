INSERT INTO casbin_rule(ptype, v0, v1, v2)
VALUES ('p', 'unauthorized', '/swagger/*', 'GET');

INSERT INTO casbin_rule(ptype, v0, v1, v2)
VALUES ('p', 'unauthorized', '/user/', 'POST');

INSERT INTO casbin_rule(ptype, v0, v1, v2)
VALUES ('p', 'authorized', '/company/', 'POST');
