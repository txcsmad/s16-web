<?php
function &get_db_connection() {
    static $conn = NULL;
    $MYSQL_HOST = "localhost";
	$MYSQL_DB_NAME = "madsite"; 
	$MYSQL_USER = "madman";
	$MYSQL_PASSWORD = "verysecurepassword";

    if (is_null($conn)) {
        try {
            $conn = new PDO("mysql:host=$MYSQL_HOST;dbname=$MYSQL_DB_NAME",
                $MYSQL_USER, $MYSQL_PASSWORD);
            $conn->setAttribute( PDO::ATTR_ERRMODE, PDO::ERRMODE_WARNING );

            return $conn;
        } catch (PDOException $e) {
            error_log($e->getMessage(), 0);
            return NULL;
        }
    }

    return $conn;
}
?>