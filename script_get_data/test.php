<?php
include('versos.php');

$consulta = new Verso("Éxodo","DHH","+8:29");
$texto=$consulta->_getResource(); // Rnc valido Test OK
echo $texto; // cedul