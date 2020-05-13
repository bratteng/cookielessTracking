# cookielessTracking

## Purpose
Proof-of-Concept that you can store a visitors tracking ID without using cookies, by utlizing their computer and local cache.

The file tracking.js is set as last-modified at 1st of January 1970, so it will be stored in cache until cache is cleared.

## Usage
```html
<script type="text/javascript" src="//example.com/tracking.js"></script>
<script type="text/javascript">
	console.log(`My tracking ID is: ${trackingID}`)
</script>
```
