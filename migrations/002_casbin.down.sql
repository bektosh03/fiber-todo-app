DELETE
FROM casbin_rule
WHERE v1 = '/swagger/*'
  AND v2 = 'GET';

DELETE
FROM casbin_rule
WHERE v1 = '/user/'
  AND v2 = 'POST';

DELETE
FROM casbin_rule
WHERE v1 = '/company/'
  AND v2 = 'POST';