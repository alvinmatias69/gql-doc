# gqlDocumentation

<details>
<summary>Table of Contents</summary>



### Queries

* [getImages](#getImages)








### Types

* [Images](#Images)

* [Data](#Data)

* [Compression](#Compression)

* [ImageQuery](#ImageQuery)



</details>


## Queries



### getImages


#### Definition
<table>
    <tr>
        <th>Return Type</th>
        <td><a href="#Images">Images</a></td>
    </tr>
    <tr>
        <th>Scalar</th>
        <td>No
    </tr>
    <tr>
        <th>Nullable</th>
        <td>No</td>
    </tr>
    <tr>
        <th>List</th>
        <td>No</td>
    </tr>
</table>

#### Parameters
Name|Type|List?|Nullable?
----|----|-----|---------
id|Int|Yes|No


#### Example

<details>
<summary>Request</summary>

```
getImages (
	id: [
		1
	]
) {
	data 
 }

```

</details>

<details>
<summary>Response</summary>

```json
{
	"getImages": {
		"data": [
			"Hello"
		]
	}
}
```

</details>








## Types


### Images

_Query to get user images by ID_


**Variant:** `Object`

#### Properties

Name|Type|Scalar|Nullable|List|Description
----|----|------|--------|----|-----------
data|String|Yes|Yes|Yes| |


### Data

_Data of given image_


**Variant:** `Input`

#### Properties

Name|Type|Scalar|Nullable|List|Description
----|----|------|--------|----|-----------
image|String|Yes|No|No|Image of given data |


### Compression

_Compression enumeration_


**Variant:** `Enum`

#### Properties

Name|Description
----|-----------
HOT|Not cold
COLD|Not hot



### ImageQuery

_ImageQuery type_


**Variant:** `Union`

#### Properties

Name|Description
----|-----------
Min|
Meta|
Complete|




