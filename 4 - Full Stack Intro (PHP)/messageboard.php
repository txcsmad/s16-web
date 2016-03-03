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

if (isset($_POST['author']) && isset($_POST['content'])) {
	$statement = $conn->prepare("INSERT INTO posts (author, content) VALUES (:author, :content)");
	$statement->execute(array('author' => $_POST['author'],
							  'content' => $_POST['content']));
	echo "Post submitted!";
}

$statement = $conn->prepare("SELECT * FROM posts");
$statement->execute();

while($result = $statement->fetch(PDO::FETCH_ASSOC)) {
	?>
	<p>
		Author: <?= $result['author'] ?>
	</p>
	<p>
		<?= $result['content'] ?>
	</p>
	<?php
}

?>

<form method="post">
	<textarea name="content"></textarea>
	<input type="text" name="author" placeholder="Your Name" />
	<button type="submit">Submit Post</button>
</form>
