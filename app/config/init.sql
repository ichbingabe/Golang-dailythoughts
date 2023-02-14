create table thoughts(
  id INT NOT NULL AUTO_INCREMENT,
  author varchar(64),
  title varchar(64),
  content VARCHAR(120),
  primary key(id)
);


insert into thoughts(title, content, author)
values
    ('Hello World', 'The obligatory Hello World Post ...', 'Pedro lençol'),
    ('Another Post', 'Yet another blog post about something exciting', 'Pascal freiatico');

    