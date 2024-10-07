// users collection
db.createCollection("users", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["id", "username", "email", "password", "createdAt"],
      properties: {
        id: {
          bsonType: "objectId",
          description: "The unique identifier for the user",
        },
        username: {
          bsonType: "string",
          description: "The username of the user",
          pattern: "^[a-zA-Z0-9_.-]+$", // Optional pattern for allowed characters
          maxLength: 50,
        },
        email: {
          bsonType: "string",
          description: "The email address of the user",
          pattern: "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$", // Email validation pattern
          maxLength: 255,
        },
        password: {
          bsonType: "string",
          description: "The password of the user",
          minLength: 8, // Optional minimum password length
        },
        createdAt: {
          bsonType: "date",
          description: "The date and time when the user was created",
        },
        updatedAt: {
          bsonType: "date",
          description: "The date and time when the user was last updated",
        },
      },
    },
  },
});

// task_lists collection
db.createCollection("taskLists", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["id", "title", "userId", "createdAt"],
      properties: {
        id: {
          bsonType: "objectId",
          description: "The unique identifier for the task list",
        },
        title: {
          bsonType: "string",
          description: "The title of the task list",
          maxLength: 255,
        },
        description: {
          bsonType: "string",
          description: "The description of the task list",
        },
        userId: {
          bsonType: "objectId",
          description: "The ID of the user who owns the task list",
          ref: "users", // Reference to the users collection
        },
        createdAt: {
          bsonType: "date",
          description: "The date and time when the task list was created",
        },
        updatedAt: {
          bsonType: "date",
          description: "The date and time when the task list was last updated",
        },
      },
    },
  },
});

// tasks collection
db.createCollection("tasks", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["id", "userId", "taskListId", "title", "createdAt"],
      properties: {
        id: {
          bsonType: "int",
          description: "The unique identifier for the task",
        },
        userId: {
          bsonType: "objectId",
          description: "The ID of the user who owns the task",
          ref: "users", // Reference to the users collection
        },
        taskListId: {
          bsonType: "objectId",
          description: "The ID of the task list the task belongs to",
          ref: "taskLists", // Reference to the taskLists collection
        },
        title: {
          bsonType: "string",
          description: "The title of the task",
          maxLength: 255,
        },
        description: {
          bsonType: "string",
          description: "The description of the task",
        },
        dueDate: {
          bsonType: "date",
          description: "The due date of the task",
        },
        completed: {
          bsonType: "bool",
          description: "Indicates whether the task is completed",
        },
        createdAt: {
          bsonType: "date",
          description: "The date and time when the task was created",
        },
        updatedAt: {
          bsonType: "date",
          description: "The date and time when the task was last updated",
        },
      },
    },
  },
});

// labels collection
db.createCollection("labels", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["id", "name", "userId", "createdAt"],
      properties: {
        id: {
          bsonType: "objectId",
          description: "The unique identifier for the label",
        },
        name: {
          bsonType: "string",
          description: "The name of the label",
          maxLength: 255,
        },
        color: {
          bsonType: "string",
          description: "The color of the label",
          maxLength: 255,
        },
        userId: {
          bsonType: "objectId",
          description: "The ID of the user who owns the label",
          ref: "users", // Reference to the users collection
        },
        createdAt: {
          bsonType: "date",
          description: "The date and time when the label was created",
        },
        updatedAt: {
          bsonType: "date",
          description: "The date and time when the label was last updated",
        },
      },
    },
  },
});

// Add indexes to improve query performance
db.users.createIndex({ username: 1 }, { unique: true });
db.users.createIndex({ email: 1 }, { unique: true });
db.taskLists.createIndex({ userId: 1 });
db.tasks.createIndex({ userId: 1 });
db.tasks.createIndex({ taskListId: 1 });
db.labels.createIndex({ userId: 1 });
