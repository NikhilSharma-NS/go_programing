there’s a powerful paradigm called “Convention over Configuration”. The concept is about taking unnecessary, time-consuming decisions out of your developer-hands, by more or less forcing you to follow a set of rules, when you develop with a specific tool or framewor



- Declare types that represent something new or unique
- Don't create alises just for readability
- Validate that a value of any type is created or used on its own
- Embed types not because i need the state but becuase we need the behavior
- if i am not thinking about behavior I am locking myself into the design that i can't grow in the future without cascading code changes
- Question types that are aliases or abstractions for an existing type
- Question type whose sole purpose is to share a common set of states