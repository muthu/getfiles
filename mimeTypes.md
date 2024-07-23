### MIME 

**MIME**: Multipurpose Internet Mail Extensions

**media** type is same as MIME type. Media type indicates the type/format of a particular file, document or group of bytes (Could be anything)

Why do we need this?

Browsers use MIME type to decide what how to process the URL and not based on the file extension. 
So it's important that web servers send the correct MIME type in the response's Content-Type header. 
If this is not correctly configured, browsers are likely to misinterpret the contents of files, sites will not work correctly, and downloaded files may be mishandled.

#### MIME type structure

`type/subtype` : (no whitespaces in between them). The type represents the general category into which the data type falls, such as video or text.
The subtype identifies the exact kind of data of the specified type the MIME type represents. For example, for the MIME type text, the subtype might be plain (plain text), html (HTML source code), or calendar (for iCalendar/.ics) files.

`type/subtype;parameter=value`: Extra parameter can be specified for example the type `text/plain;charset=UTF-8` represents the mime type text and subtype plain text with character set used as UTF-8

#### MIME Types

There are two classes of MIME types:

1. Discrete: represents a single file or medium. Example: single text file or single music file
2. Multi-part: represents a composite document whose group of components each of which can have its own MIME type or a multipart type may encapsulate multiple files being sent together in one transaction. For example, multipart MIME types are used when attaching multiple files to an email.

For actual examples of these classes visit [here](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types)
