# HTTP/2 Header Compression

In HTTP/2, one of the significant improvements over HTTP/1.x is **header compression**. The primary goal of this feature is to reduce the overhead of repetitive HTTP headers, which is especially important in scenarios where many requests share the same set of headers (for example, cookies, user-agent, host, etc.). This results in more efficient use of network bandwidth and improved performance.

### **Overview of HTTP/2 Header Compression**

HTTP/2 uses a specialized compression algorithm called **HPACK** (HTTP/2 Header Compression) to handle header compression. HPACK is designed to be efficient, minimize memory usage, and ensure security, while also being resistant to certain types of attacks (like **denial-of-service (DoS)** attacks) that could exploit header compression.

### **Why Header Compression Is Needed in HTTP/2**

In HTTP/1.x, every HTTP request and response contains a set of headers. These headers are sent in plain text and are often very similar for every request (e.g., the `User-Agent` or `Host` header). In high-latency networks or applications that require a lot of requests (e.g., a web page with many resources like images, scripts, and stylesheets), this redundancy can lead to significant performance overhead. The repeated transmission of the same headers increases the size of each request, adding to the latency and consuming more bandwidth.

In HTTP/2, multiple requests and responses can share the same underlying TCP connection. Therefore, header compression is applied to reduce the overhead of sending the same headers with each new request or response.

### **HPACK: The HTTP/2 Header Compression Format**

HPACK is the compression format used in HTTP/2 to efficiently encode and decode HTTP headers. HPACK has been specifically designed to work with HTTP/2's features, such as multiplexing and stream prioritization, and addresses the following goals:

- **Minimize redundancy**: Many HTTP headers are sent repeatedly (e.g., `User-Agent`, `Content-Type`, etc.). HPACK reduces the size of these headers by encoding them efficiently.
- **Allow dynamic header table**: HPACK uses a dynamic table to store previously sent headers. This allows new headers to be encoded by referencing a previously transmitted header, reducing the need to send full header values every time.
- **Protect against compression attacks**: HTTP/2’s HPACK uses techniques to prevent certain types of attacks that could exploit header compression (e.g., **denial-of-service** attacks due to the large memory requirements for decompressing headers).

### **How HPACK Works**

HPACK's header compression works in two main ways:

1. **Static Table**:  
   The **static table** is a predefined list of common HTTP headers that are often seen in most HTTP requests and responses. These headers are assigned an index, and when they appear in a request or response, they can be referred to by index instead of being sent in full text.

   Examples of headers in the static table include:
   - `:method`
   - `:scheme`
   - `:status`
   - `accept-encoding`
   - `content-type`
   - `user-agent`

   The static table is fixed and shared by both the client and server, allowing headers to be encoded using smaller references rather than full strings.

2. **Dynamic Table**:  
   In addition to the static table, HPACK also includes a **dynamic table**. The dynamic table allows new headers that have not yet been seen in the session to be added to the table as the communication progresses. This dynamic table is updated during the communication and is specific to each HTTP/2 connection.

   - When a header is added to the dynamic table, it is assigned an index.
   - Future headers can refer to previous headers from the dynamic table by their index instead of sending the entire header again.
   - Each header in the dynamic table has a **name** and **value**. The dynamic table grows as headers are sent.

### **Encoding and Decoding of Headers with HPACK**

**Encoding**: When a new header needs to be sent, HPACK uses one of the following methods:
1. **Literal Encoding (with or without Indexing)**:
   - If the header is new or not part of the static or dynamic table, it can be sent as a literal header field with its name and value.
   - If the header has appeared before, the encoder may choose to reference it using an index from the static or dynamic table.

2. **Indexed Header**:
   - If a header already exists in the static or dynamic table, it can be sent by index. This reference is much smaller than sending the full header name and value.

3. **Incremental Indexing**:  
   - HPACK supports **incremental indexing** for headers that appear frequently. Once a header is sent, it is added to the dynamic table, and future occurrences of that header can be sent using its index, avoiding repeated transmission of the same value.

**Decoding**: The decoder at the receiving end works in reverse:
1. It either looks up the header in the static or dynamic table by index.
2. If the header is new, it is added to the dynamic table and indexed for future use.

### **HPACK Header Format**

Each HPACK header is encoded with a series of integer values that specify the operation to be performed (such as adding a header to the dynamic table or referring to an existing one).

For example:
- **Literal Header Field with Indexing**: This is used when a header is sent by index from the static or dynamic table.
- **Literal Header Field without Indexing**: When a header is new or not in the table, it is sent in full (name and value).
- **Indexed Header Field**: This refers to a header already in the static or dynamic table by index.

### **HPACK Example**

Let’s consider an example HTTP/2 request:

#### Request Without Compression (HTTP/1.x Style):
```
GET /index.html HTTP/2.0
Host: example.com
User-Agent: Mozilla/5.0
Accept-Encoding: gzip, deflate
```

Without header compression, this request would contain a lot of redundant information, especially headers like `Host` and `User-Agent`.

#### Request with HPACK Compression (HTTP/2):
- **Static Table Indexing**: `Host` and `User-Agent` may already be present in the static table, so their values can be referenced by an index.
- **Dynamic Table Usage**: If `Accept-Encoding: gzip, deflate` has not been seen before in the current session, it will be added to the dynamic table for future use, and subsequent requests can reference it by index.

### **Security Considerations of HPACK**

HPACK was designed with security in mind, addressing potential vulnerabilities such as **compression-based DoS attacks**. For example, without proper protections, attackers could craft headers that cause excessive memory usage when decompressed.

- **HPACK is resistant to such attacks** because it uses a **bounded memory** approach. The size of the dynamic table is limited, and if the table grows too large, it must be cleared.
- **HPACK maintains a separate encoding and decoding process** for both client and server, ensuring that they stay synchronized while preventing attackers from manipulating the compression state.

### **Impact of Header Compression on Performance**

1. **Reduced Latency**:  
   By compressing headers, HTTP/2 significantly reduces the size of requests and responses, resulting in faster transmission times and reduced network congestion. This is particularly important for mobile devices or networks with limited bandwidth.

2. **Bandwidth Efficiency**:  
   HTTP/2’s header compression minimizes the amount of data transferred, saving bandwidth and improving the performance of web applications, especially those with many repeated headers (like cookies or custom headers).

3. **Improved Multiple Requests Handling**:  
   HTTP/2's multiplexing combined with header compression allows browsers to send multiple requests and responses over a single connection without re-sending the same headers. This reduces the number of round trips and further improves performance.

### **Conclusion**

HTTP/2 header compression (via HPACK) is a critical feature for reducing the overhead of redundant headers, especially in applications with high-frequency requests and responses. By efficiently encoding headers with static and dynamic tables, HTTP/2 significantly improves network performance and reduces latency. HPACK offers a balance between compression efficiency and security, ensuring that applications can send data more efficiently without sacrificing safety.

Header compression in HTTP/2 plays an essential role in speeding up web traffic, making it one of the key enhancements over HTTP/1.x and contributing to the overall success of the protocol.