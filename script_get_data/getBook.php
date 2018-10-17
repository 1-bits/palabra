<?php
include('conection.php');
include('versos.php');

function guardarData($id_version, $id_libro, $capitulo, $verso, $texto){
    try {
        $sql = "INSERT INTO versos( version,libro,capitulo,verso,texto)
                    VALUES (?,?,?,?,?)";
        $conection = new Conection();
        $conection->pdo()->prepare($sql)->execute(
            array(
                $id_version,
                $id_libro,
                $capitulo,
                $verso,
                $texto,
            )
        );
    } catch (Exception $e) {
        die($e->getMessage());
    }
}
$resultversion = array();
$resultbook = array();
try {
    $conection = new Conection();
    $stm = $conection->pdo()->prepare("SELECT * FROM version");
    $stm->execute();
    foreach ($stm->fetchAll(PDO::FETCH_OBJ) as $r) {
        $obj = new stdClass;
        $obj->id = $r->id;
        $obj->version = $r->version;
        array_push($resultversion, $obj);
    }
} catch (Exception $e) {
    die("error en el SQL :" . $e);
}
try {
    $conection = new Conection();
    $stm = $conection->pdo()->prepare("SELECT * FROM libros");
    $stm->execute();
    foreach ($stm->fetchAll(PDO::FETCH_OBJ) as $r) {
        $obj1 = new stdClass;
        $obj1->id = $r->id;
        $obj1->nombre = $r->nombre;
        array_push($resultbook, $obj1);
    }
} catch (Exception $e) {
    die("error en el SQL :" . $e);
}

foreach ($resultversion as $obj_version) {
    foreach ($resultbook as $obj_book) {
        printf("Iniciamdo con el Libro de " . $obj_book->nombre . ", trabajamos con la version " .
            $obj_version->version . " \n");
        $cap = 1;
        while ($cap >= 1) {
            $ver = 1;
            while ($ver >= 1) {
                $consulta = new Verso($obj_book->nombre, $obj_version->version, "+" . strval($cap) . ":" . strval($ver));
                $texto = $consulta->_getResource();
                if ($texto == "") {
                    if ($ver == 1) {
                        $cap = -1;
                    }
                    $ver = -1;
                }
                if ($ver != -1) {
                    guardarData($obj_version->id, $obj_book->id, $cap, $ver, $texto);
                }
                $ver++;
            }
            $cap++;
        }
    }
}
   