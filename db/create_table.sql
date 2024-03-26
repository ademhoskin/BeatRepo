create table if not exists BeatZip (
  zipId serial primary key not null,
  zipName varchar(255) not null,
  zipSize bigint not null,
  S3Link text not null,
  dateAdded date not null
);
