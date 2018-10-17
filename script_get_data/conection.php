<?php
include dirname(__FILE__).'\config.php';

class Conection {
  
    private $connection=null;
    public function __construct(){
          global $config;
        try {
            $dbsystem=$config['db_system'];
            $host=$config['db_host'];
            $dbname=$config['db_name'];
            $charset=";charset=utf8";
            $dsn=$dbsystem.':host='.$host.';dbname='.$dbname.$charset;
            $username=$config['db_username'];
            $passwd=$config['db_password'];
            $this->connection = new PDO($dsn, $username, $passwd);
            $this->connection->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
        } catch (PDOException $e) {
            print "¡Error!: " . $e->getMessage() . "<br/>";
            die();
        }
    }

    public function pdo() {
        return $this->connection;
    }

}