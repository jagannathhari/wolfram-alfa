package main

import rl "github.com/gen2brain/raylib-go/raylib"


const (
     W =  800  
     H = 800
     TITLE =  "Wolfram Celluar Automation"
     SIZE =  10
     ROWS = H/SIZE
     COLS = W/SIZE
     FPS =  24
)


func GetState(neiougher int32,rule int32) bool{
    return (rule & (1 << (8 - neiougher - 1))) > 0;
}

func RandomFill(array []bool,len int) {
    state := [2]bool{false,true}
    for i:= 1; i <= len; i++ {
        array[i] = state[rl.GetRandomValue(0, 1)]
    }
}

func FillWith[T any](array[]T , value T){
    for i:= range array{
        array[i] = value
    }
}

func UpdateState(current_state[]bool, next_state[]bool,cols int32,rule int32) {
    var neioughors int32 
    for i := 1; i <= COLS; i++{
        neioughors = 0;
        if (current_state[i - 1]){
            neioughors = neioughors | 4;
        }

        if (current_state[i]){
            neioughors = neioughors | 2;
        }

        if (current_state[i + 1]){
            neioughors = neioughors | 1;
        }

        next_state[i] = GetState(neioughors, rule);
    }
}

func main() {
	rl.InitWindow(W, H,TITLE )
	defer rl.CloseWindow()

	rl.SetTargetFPS(FPS)

    var rule int32 = 195

    grid := make([][]bool, ROWS)
    state := make([]bool, COLS+2)
    temp := make([]bool, ROWS)

    for i:=0 ; i< ROWS ; i++ {
        grid[i] = make([]bool,COLS+2)
    }

    grid[ROWS - 1][COLS] = true;

	for !rl.WindowShouldClose() {
        // Get random rule and fill random location
        if (rl.IsKeyReleased(rl.KeyR)) {
            rule = rl.GetRandomValue(0, 254);
            RandomFill(grid[ROWS - 1], COLS);
        }

        if (rl.IsKeyReleased(rl.KeyJ) && rule < 254) {
            rule++;
            FillWith(grid[ROWS-1],false)
            grid[ROWS - 1][COLS] = true;
        }
       
        if (rl.IsKeyReleased(rl.KeyK) && rule > 0) {

            FillWith(grid[ROWS-1],false)
            grid[ROWS - 1][COLS] = true;
            rule--;
        }


		rl.BeginDrawing()

        for i:= 0; i < ROWS; i++{
            for j := 1; j <= COLS; j++ {
                if (grid[i][j]) {
                    rl.DrawRectangle(int32((j - 1) * SIZE), int32(i * SIZE), SIZE, SIZE, rl.Red);
                }
            }
        }
		rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()
        UpdateState(grid[ROWS - 1], state, COLS, rule);
        temp = grid[0];
        for i := 1; i < ROWS; i++ {
            grid[i - 1] = grid[i];
        }
        grid[ROWS - 1] = state;
        state = temp;
	}
}
