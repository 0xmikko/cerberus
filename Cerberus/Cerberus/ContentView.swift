//
//  ContentView.swift
//  Cerberus
//
//  Created by Mikhail Lazarev on 23.10.2019.
//  Copyright © 2019 Mikhail Lazarev. All rights reserved.
//

import SwiftUI

struct ContentView: View {
    
    @State var address2 : String = ""
    
    var body: some View {
        VStack {
            Text("Welcome to Cerberus")
            TextField<Text>(/*@START_MENU_TOKEN@*/"Placeholder"/*@END_MENU_TOKEN@*/, text: $address2, onEditingChanged: { n in
                print(n)
                
            })
                .font(/*@START_MENU_TOKEN@*/.title/*@END_MENU_TOKEN@*/)
            
            Button(action: /*@START_MENU_TOKEN@*/{}/*@END_MENU_TOKEN@*/) {
                Text(/*@START_MENU_TOKEN@*/"Button"/*@END_MENU_TOKEN@*/)
                    .background(Color.red)
                    .cornerRadius(10)
                    .padding(.horizontal)
                    .foregroundColor(Color.white)
                    
            }
        }
        .padding(/*@START_MENU_TOKEN@*/.all/*@END_MENU_TOKEN@*/)
        
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
