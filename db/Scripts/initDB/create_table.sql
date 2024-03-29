create table if not exists BeatZip (
  zipId serial primary key not null,
  zipName varchar(255) not null,
  zipSize bigint not null,
  s3Link text not null,
  dateAdded date not null
);
