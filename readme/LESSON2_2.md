# Lesson 2.2: Practice problems

## Practice exercises

### 1. Guess the Number Game  
**Problem:**  
Create a command-line game where the program randomly selects a number between 1 and 100, and the user has to guess what it is.
After each guess, provide feedback:  
- "Too low!" if the guess is below the number  
- "Too high!" if the guess is above the number  
- "Correct!" when the guess is right  

The game should keep running until the correct number is guessed.  
**Skills practiced:** loops, conditionals, user input, randomness

---

### 2. Two Sum Problem  
**Problem:**  
Write a function that takes a slice of integers and a target integer, and returns the **indices** of the two numbers that add up to the target.

**Constraints:**  
- Do not use the same element twice  
- Exactly one solution is guaranteed  
- You can return the result in any order  

**Example:**  
```go
nums := []int{2, 7, 11, 15}
target := 9
// Output: [0, 1] because nums[0] + nums[1] == 9
```

### 3. Auction REST API  
**Problem:**  
Create a basic HTTP REST API that simulates an auction website. You don’t need any frontend, just API routes that can be tested with curl, Postman, or a browser. Use only in-memory storage (e.g., a global slice or map).

Each product should have:
- `Name` (string)
- `Description` (string)
- `Price` (float64)
- `IsSold` (bool)

**Required API endpoints:**
1. **Add a product** – `POST /products`  
   - Request body: JSON `{ "name": "Laptop", "description": "Gaming laptop", "price": 999.99 }`
2. **Remove a product** – `DELETE /products/{name}`  
3. **Sell a product** – `POST /products/{name}/sell`  
4. **List all products** – `GET /products`  
5. **List only sold products** – `GET /products?sold=true`  
6. **General product view** – `GET /products`  
   - Response: list of all products showing only `name` and `price`
7. **Detailed product view** – `GET /products/{name}`  
   - Response: includes `name`, `price`, `description`, `isSold`

**Skills practiced:**  
- Creating an HTTP server with `net/http`  
- REST API design  
- Structs, slices, basic state management  
- JSON marshaling/unmarshaling  
- Route handling and URL parameters
