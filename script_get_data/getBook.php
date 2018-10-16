<?php
include('conection.php');
include('versos.php');


$result = array();
try{
    $conection= new Conection();
    $stm = $conection->pdo()->prepare("SELECT * FROM libros");
    $stm->execute();
    foreach($stm->fetchAll(PDO::FETCH_OBJ) as $r)
    {   
        array_push($result, $r->id,$r->nombre);
    }
    }catch(Exception $e){
        die("error en el SQL :".$e);
    } 

    var_dump( $result);
    $consulta = new Verso("mateo","LBLA","+1:26");
    $texto=$consulta->_getResource();
    var_dump($texto);