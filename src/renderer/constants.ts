export const buildSqlitePath = (basePath) =>
  `${basePath}/.kobo/KoboReader.sqlite`

export const bookQuery = `
SELECT DISTINCT
  b.VolumeID, c.Title, c.Attribution, c.___PercentRead
FROM
  Content c
INNER JOIN
  Bookmark b on c.ContentID = b.VolumeID
WHERE
  c.ContentType = 6 AND c.MimeType = 'application/x-kobo-epub+zip'
`
