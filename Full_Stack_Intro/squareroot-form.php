<?php

if (ctype_digit($_GET['value'])) {
	$value = intval($_GET['value']);
	
	$result = "The square root of {$_GET['value']} is " . sqrt($value);
} else {
	$result = "Input invalid or not specified.";
}

?>
<h1>Square Root-er</h1>
<h2>Your Square Root or Your Money Back</h2>
<span>

<?= $result ?>

</span>

<form method="get" action="squareroot-form.php">
	<input type="text" name="value" />
	<button type="submit">Get Your Square Root</button>
</form>