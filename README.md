# Dota2API
<p>This is an API developed in <a href=https://go.dev/>GO</a> based in the videogame Dota 2</p>
<h3>Used Technologies</h3>
<li><a href=https://github.com/gin-gonic/gin>Gin Gonic</a></li>
<li><a href=https://gorm.io/>Gorm</a></li>
<li><a href=https://www.postgresql.org//>PostgreSQL</a></li>

<h3>Getting Started</h3>
<p>To configure the server, go to the main file, inside the main function, 
in the call to the Run method, you pass a string with the address where 
you want the server as a parameter.</p>
<p>The SetupDB file contains the configuration of the database with which the API will work:
<li>Host</li>
<li>User</li>
<li>Password</li>
<li>Dbname</li>
<li>Port</li>
</p>


<h1>Endpoints</h1>

<h3>Hero Endpoints</h3>
<li>Get all heroes in the database</li>
<h6>URL:http://api/v1/heroes</h6>
<h6>Method:GET</h6>

<li>Get heroes by the attribute</li>
<h6>URL:http://api/v1/heroes/attribute/:atributte</h6>
<h6>Method:GET</h6>

<li>Get Hero by name</li>
<h6>URL:http://api/v1/heroes/:name</h6>
<h6>Method:GET</h6>

<li>Post one hero in the database</li>
<h6>URL:http://api/v1/heroes</h6>
<h6>Method:POST</h6>

<li>Delete hero by id</li>
<h6>URL:http://api/v1/heroes/:id</h6>
<h6>Method:DELETE</h6>

<h3>Item Endpoints</h3>
<li>Get all items in the database</li>
<h6>URL:http://api/v1/heroes</h6>
<h6>Method:GET</h6>

<li>Get Item by name</li>
<h6>URL:http://api/v1/heroes/:name</h6>
<h6>Method:GET</h6>

<li>Post one item in the database</li>
<h6>URL:http://api/v1/heroes</h6>
<h6>Method:POST</h6>

<li>Delete item by id</li>
<h6>URL:http://api/v1/heroes/:id</h6>
<h6>Method:DELETE</h6>

<h3>Skill Endpoints</h3>
<li>Get all skills in the database</li>
<h6>URL:http://api/v1/heroes</h6>
<h6>Method:GET</h6>

<li>Get Skill by name</li>
<h6>URL:http://api/v1/heroes/:name</h6>
<h6>Method:GET</h6>

<li>Post one skill in the database</li>
<h6>URL:http://api/v1/heroes</h6>
<h6>Method:POST</h6>

<li>Delete skill by id</li>
<h6>URL:http://api/v1/heroes/:id</h6>
<h6>Method:DELETE</h6>


