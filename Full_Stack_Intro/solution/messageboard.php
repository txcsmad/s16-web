<?php
require 'db.php';

$conn = get_db_connection();
if ($conn === false) {
	?>
	<h1>Error connecting to database</h1>
	<?php
	die();
}

/* GET READY */
/* PUT YOUR CODE HERE */
?>

<form method="post">
	<textarea name="content"></textarea>
	<input type="text" name="author" placeholder="Your Name" />
	<button type="submit">Submit Post</button>
</form>
