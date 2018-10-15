<?php
include('versos.php');
$consulta = new Verso("mateo","+1:1","LBLA");
$consulta->queryDoc() ; // Rnc valido Test OK
//echo $consulta->queryDoc(); // cedul