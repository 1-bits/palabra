DROP DATABASE IF EXISTS palabradb;
CREATE DATABASE palabradb;
USE palabradb;

CREATE TABLE `version` (
  `id` int(3) unsigned zerofill NOT NULL AUTO_INCREMENT,
  `version` text NOT NULL,
  `descripcion` text NOT NULL,
  `idioma` text NOT NULL,
  PRIMARY KEY (`id`)
);
insert into version(id,version,descripcion,idioma)  value(1,'DHH','Dios Habla Hoy (DHH)','');
insert into version(id,version,descripcion,idioma)  value(2,'JBS','Jubilee Bible 2000 (Spanish) (JBS)','');
insert into version(id,version,descripcion,idioma)  value(3,'NBLHN','Nueva Biblia Latinoamericana de Hoy (NBLH)','');
insert into version(id,version,descripcion,idioma)  value(4,'NBV','Nueva Biblia Viva (NBV)','');
insert into version(id,version,descripcion,idioma)  value(5,'NTV','Nueva Traducción Viviente (NTV)','');
insert into version(id,version,descripcion,idioma)  value(6,'NVI','Nueva Versión Internacional (NVI)','');
insert into version(id,version,descripcion,idioma)  value(7,'CST','Nueva Versión Internacional (Castilian) (CST)','');
insert into version(id,version,descripcion,idioma)  value(8,'PDT','Palabra de Dios para Todos (PDT)','');
insert into version(id,version,descripcion,idioma)  value(9,'BLP','La Palabra (España) (BLP)','');
insert into version(id,version,descripcion,idioma)  value(10,'BLPH','La Palabra (Hispanoamérica) (BLPH)','');
insert into version(id,version,descripcion,idioma)  value(11,'RVA-2015','Reina Valera Actualizada (RVA-2015)','');
insert into version(id,version,descripcion,idioma)  value(12,'RVC','Reina Valera Contemporánea (RVC)','');
insert into version(id,version,descripcion,idioma)  value(13,'RVR1960','Reina-Valera 1960 (RVR1960)','');
insert into version(id,version,descripcion,idioma)  value(14,'RVR1977','Reina Valera Revisada (RVR1977)','');
insert into version(id,version,descripcion,idioma)  value(15,'RVR1995','Reina-Valera 1995 (RVR1995)','');
insert into version(id,version,descripcion,idioma)  value(16,'RVA','Reina-Valera Antigua (RVA)','');
insert into version(id,version,descripcion,idioma)  value(17,'SRV-BRG','Spanish Blue Red and Gold Letter Edition (SRV-BRG)','');
insert into version(id,version,descripcion,idioma)  value(18,'TLA','Traducción en lenguaje actual (TLA)','');

CREATE TABLE `libros` (
  `id` int(11) UNSIGNED NOT NULL,
  `nombre` varchar(50) COLLATE latin1_spanish_ci NOT NULL,
 PRIMARY KEY (`id`)
);
insert into libros(id,nombre)  value(1,'Génesis');
insert into libros(id,nombre)  value(2,'Éxodo');
insert into libros(id,nombre)  value(3,'Levítico');
insert into libros(id,nombre)  value(4,'Números');
insert into libros(id,nombre)  value(5,'Deuteronomio');
insert into libros(id,nombre)  value(6,'Josué');
insert into libros(id,nombre)  value(7,'Jueces');
insert into libros(id,nombre)  value(8,'Rut');
insert into libros(id,nombre)  value(9,'1 Samuel');
insert into libros(id,nombre)  value(10,'2 Samuel');
insert into libros(id,nombre)  value(11,'1 Reyes');
insert into libros(id,nombre)  value(12,'2 Reyes');
insert into libros(id,nombre)  value(13,'1 Crónicas');
insert into libros(id,nombre)  value(14,'2 Crónicas');
insert into libros(id,nombre)  value(15,'Esdras');
insert into libros(id,nombre)  value(16,'Nehemías');
insert into libros(id,nombre)  value(17,'Ester');
insert into libros(id,nombre)  value(18,'Job');
insert into libros(id,nombre)  value(19,'Salmos');
insert into libros(id,nombre)  value(20,'Proverbios');
insert into libros(id,nombre)  value(21,'Eclesiastés');
insert into libros(id,nombre)  value(22,'Cantares');
insert into libros(id,nombre)  value(23,'Isaías');
insert into libros(id,nombre)  value(24,'Jeremías');
insert into libros(id,nombre)  value(25,'Lamentaciones');
insert into libros(id,nombre)  value(26,'Ezequiel');
insert into libros(id,nombre)  value(27,'Daniel');
insert into libros(id,nombre)  value(28,'Oseas');
insert into libros(id,nombre)  value(29,'Joel');
insert into libros(id,nombre)  value(30,'Amós');
insert into libros(id,nombre)  value(31,'Abdías');
insert into libros(id,nombre)  value(32,'Jonás');
insert into libros(id,nombre)  value(33,'Miqueas');
insert into libros(id,nombre)  value(34,'Nahúm');
insert into libros(id,nombre)  value(35,'Habacuc');
insert into libros(id,nombre)  value(36,'Sofonías');
insert into libros(id,nombre)  value(37,'Hageo');
insert into libros(id,nombre)  value(38,'Zacarías');
insert into libros(id,nombre)  value(39,'Malaquías');
insert into libros(id,nombre)  value(40,'Mateo');
insert into libros(id,nombre)  value(41,'Marcos');
insert into libros(id,nombre)  value(42,'Lucas');
insert into libros(id,nombre)  value(43,'Juan');
insert into libros(id,nombre)  value(44,'Hechos');
insert into libros(id,nombre)  value(45,'Romanos');
insert into libros(id,nombre)  value(46,'1 Corintios');
insert into libros(id,nombre)  value(47,'2 Corintios');
insert into libros(id,nombre)  value(48,'Gálatas');
insert into libros(id,nombre)  value(49,'Efesios');
insert into libros(id,nombre)  value(50,'Filipenses');
insert into libros(id,nombre)  value(51,'Colosenses');
insert into libros(id,nombre)  value(52,'1 Tesalonicenses');
insert into libros(id,nombre)  value(53,'2 Tesalonicenses');
insert into libros(id,nombre)  value(54,'1 Timoteo');
insert into libros(id,nombre)  value(55,'2 Timoteo');
insert into libros(id,nombre)  value(56,'Tito');
insert into libros(id,nombre)  value(57,'Filemón');
insert into libros(id,nombre)  value(58,'Hebreos');
insert into libros(id,nombre)  value(59,'Santiago');
insert into libros(id,nombre)  value(60,'1 Pedro');
insert into libros(id,nombre)  value(61,'2 Pedro');
insert into libros(id,nombre)  value(62,'1 Juan');
insert into libros(id,nombre)  value(63,'2 Juan');
insert into libros(id,nombre)  value(64,'3 Juan');
insert into libros(id,nombre)  value(65,'Judas');
insert into libros(id,nombre)  value(66,'Apocalipsis');

CREATE TABLE `versos` (
  `id` int(11) UNSIGNED NOT NULL,
  `version` varchar(50) COLLATE latin1_spanish_ci NOT NULL,
  `libro` varchar(10) COLLATE latin1_spanish_ci NOT NULL,
  `capitulo` varchar(1) COLLATE latin1_spanish_ci NOT NULL,
  `verso` varchar(45) COLLATE latin1_spanish_ci NOT NULL,
  `texto` text NOT NULL,
  PRIMARY KEY (`id`)
);

